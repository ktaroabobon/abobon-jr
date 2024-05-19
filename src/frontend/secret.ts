import {dotenv} from "../../deps.ts"

dotenv.configSync({
  export: true,
  path: "../../.env",
})

const config = Deno.env.toObject();

export const Secret = {
  DISCORD_TOKEN: config["DISCORD_BOT_TOKEN"],
  DISCORD_CLIENT_ID: config["DISCORD_CLIENT_ID"],
  DISCORD_GUILD_ID: config["DISCORD_GUILD_ID"],
  CINII_APP_ID: config["CINII_APP_ID"],
}