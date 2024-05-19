// Import DenoDB components
import { DataTypes, Model } from "../../../deps.ts";

// Define the Paper model
class Paper extends Model {
  static table = "papers";
  static fields = {
    id: { primaryKey: true, autoIncrement: true },
    title: DataTypes.STRING,
    authors: DataTypes.STRING,
    publicationDate: DataTypes.DATE,
    publisher: DataTypes.STRING,
    publicationName: DataTypes.STRING,
    doi: DataTypes.STRING,
    naid: DataTypes.STRING,
    url: DataTypes.STRING,
  };
}
