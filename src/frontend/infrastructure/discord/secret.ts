import {dotenv} from "./deps.ts"

dotenv.configSync({
  export: true,
  path: "../../../../.env",
})

export const Secret = {
  DISCORD_TOKEN: Deno.env.get("DISCORD_BOT_TOKEN")!,
  DISCORD_CLIENT_ID: Deno.env.get("DISCORD_CLIENT_ID")!,
  DISCORD_GUILD_ID: Deno.env.get("DISCORD_GUILD_ID")!,
  CINII_APP_ID: Deno.env.get("CINII_APP_ID")!,
}