REJOHA 2022 - Responsible News Recommender Systems
==================================================

https://hack.opendata.ch/project/893

## Step 1:

Based on the aggregated popularity data (available for certain content on the
SRF APIs, for example), the first step would be to try to find out which
"categories" of content are over- or under-represented in the most popular
content.

### Top 200 videos

Generate a list of top 200 videos:

  export SRG_OAUTH_TOKEN=xyz
  go run cmd/top200/main.go > /tmp/top-200.tsv

Sample output:

  Popularity	Category	Title	Id
  19483	Sport	Doppelter Richarlison schiesst Brasilien mit Traumtor zum Sieg	516a880f-366f-43bb-9075-daccf1eddc1c
  14244	News	Tagesschau vom 24.11.2022: Hauptausgabe	60aa9e97-cfb1-48bb-951b-363107cb46a0
  8059	News	10 vor 10 vom 24.11.2022	17c59931-8c51-4509-a3dd-6e67c3416b5f
  7630	Sport	Schweizerinnen kämpfen um EM-Gold	fe0c0275-dbc6-4d3d-8441-9ebedf52df7c
  6259	Kultur	Mythos Matter	6702d57e-2f65-424d-8a0a-176a65b5e1f0
  ...

### Top 200 videos aggregated by category

Aggregate number of videos for every category:

  export SRG_OAUTH_TOKEN=xyz
  go run cmd/catpop200/main.go > /tmp/catpop-200.tsv

Sample output:

  Popularity	Category
  66303	Sport
  46909	News
  28289	Unterhaltung
  13639	Film
  10798	Dok & Reportagen
  10666	Kultur
  9730	Wissen & Ratgeber


### Top 200 videos aggregated by views

Aggregate number of views for every category:

  export SRG_OAUTH_TOKEN=xyz
  go run cmd/catcount200/main.go > /tmp/catcount-200.tsv

Sample output:

  Sum	Category
  38	Sport
  31	News
  20	Unterhaltung
  19	Dok & Reportagen
  18	Film
  17	Serien

## Step 2: tbd
