import {createBot, Intents, startBot, CreateSlashApplicationCommand, InteractionResponseTypes} from "./deps.ts";
import {Secret} from "./secret.ts";
import {SearchPapersUseCase} from "../../../backend/application/usecases/searchPapers.usecase.ts";
import {PaperSearchServiceImpl} from "../../../backend/infrastructure/cinii/paperSearch.service.impl.ts";
import {PaperRepositoryImpl} from "../../../backend/infrastructure/sqlite/paper.repository.impl.ts";
import {FetchLatestRssItemsUseCase} from "../../../backend/application/usecases/fetchLatestRssItems.usecase.ts";
import {RssItemRepositoryImpl} from "../../../backend/infrastructure/rss/rssItem.repository.impl.ts";

const bot = createBot({
  token: Secret.DISCORD_TOKEN,
  intents: Intents.Guilds | Intents.GuildMessages | Intents.MessageContent,
  events: {
    ready: (_bot, payload) => {
      console.log(`${payload.user.username} の準備が完了しました！`);

      // // 使用しないコマンドを削除
      // const guild = _bot.guilds.get(Secret.DISCORD_GUILD_ID);
      // if (guild) {
      //   guild.commands.set([])
      //     .then(console.log)
      //     .catch(console.error);
      // }
    },
  },
});

const connectionString = "mysql---"
const paperRepository = new PaperRepositoryImpl(connectionString);
const rssItemRepository = new RssItemRepositoryImpl();
const RSS_FEED_URL = "https://ktaroabobon.github.io/index.xml";

const pingCommand: CreateSlashApplicationCommand = {
  name: "ping",
  description: "Botの応答速度をチェックします",
};

const thesisCommand: CreateSlashApplicationCommand = {
  name: "thesis",
  description: "論文を検索します",
  options: [
    {
      name: "keywords",
      description: "検索キーワード（複数指定可）",
      type: 3, // STRING
      required: true,
    },
  ],
};

const abobonArticlesCommand: CreateSlashApplicationCommand = {
  name: "abobon-articles",
  description: "Abobonのサイトの最新記事を取得します",
};

await bot.helpers.upsertGuildApplicationCommands(Secret.DISCORD_GUILD_ID, [pingCommand, thesisCommand, abobonArticlesCommand]);

bot.events.messageCreate = (b, message) => {
  if (message.content === "!ping") {
    console.log("Received ping command");
    b.helpers.sendMessage(message.channelId, {
      content: "Pong!",
    });
  }
};

bot.events.interactionCreate = async (b, interaction) => {
  switch (interaction.data?.name) {
    case "ping": {
      console.log("Received /ping command");
      await b.helpers.sendInteractionResponse(interaction.id, interaction.token, {
        type: InteractionResponseTypes.ChannelMessageWithSource,
        data: {
          content: "Pong!",
        },
      });
      break;
    }
    case "thesis": {
      console.log("Received /thesis command");

      const keywordsOption = interaction.data?.options?.find((option) => option.name === "keywords");
      if (!keywordsOption || !keywordsOption.value) {
        await b.helpers.sendInteractionResponse(interaction.id, interaction.token, {
          type: InteractionResponseTypes.ChannelMessageWithSource,
          data: {
            content: "キーワードを指定する必要があります。例: /thesis keyword1 keyword2",
          },
        });
        break;
      }

      const keywords = (keywordsOption.value as string).split(" ");
      console.log(`Searching for papers with keywords: ${keywords}`);

      await b.helpers.sendInteractionResponse(interaction.id, interaction.token, {
        type: InteractionResponseTypes.DeferredChannelMessageWithSource,
      });

      await papaerRepository.connect();
      await new SearchPapersUseCase(
        new PaperSearchServiceImpl(Secret.CINII_APP_ID),
        paperRepository
      ).execute(keywords.join(" "));
      await paperRepository.disconnect();

      console.log(`Found ${papers.length} papers`);

      const response = papers
        .map(
          (paper) =>
            `タイトル: ${paper.title}\n著者: ${paper.authors}\nURL: ${paper.url}`
        )
        .join("\n\n");

      await b.helpers.editOriginalInteractionResponse(interaction.token, {
        content: response,
      });
      break;
    }
    case "abobon-articles": {
      console.log("Received /abobon-articles command");

      console.log(`Fetching Abobon articles from: ${RSS_FEED_URL}`);

      await b.helpers.sendInteractionResponse(interaction.id, interaction.token, {
        type: InteractionResponseTypes.DeferredChannelMessageWithSource,
      });

      const items = await new FetchLatestRssItemsUseCase(
        rssItemRepository
      ).execute(RSS_FEED_URL, 5);

      console.log(`Fetched ${items.length} latest articles`);

      const response = items
        .map((item) => `タイトル: ${item.title.value}\nリンク: ${item.link.href}`)
        .join("\n\n");

      await b.helpers.editOriginalInteractionResponse(interaction.token, {
        content: response,
      });
      break;
    }
    default: {
      break;
    }
  }
};

await startBot(bot);