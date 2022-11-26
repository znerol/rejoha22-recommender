REJOHA 2022 - Responsible News Recommender Systems
==================================================

https://hack.opendata.ch/project/893

## Step 1: Scrape data from SRG most-clicked video endpoint

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

## Step 2: Boost under-represented content in most-clicked video list

The top hits of the most clicked videos are often from the same category
(sports, news). A simple way to boost content from less represented categories
to a better rank is to normalize the views within categories. Each entry is
given a score between 0 and 1, and the list is ordered by that score instead of
by the number of clicks.

### Top 200 videos with clicks normalized within each category

The most basic way to compute the score is to simply normalize the number of
clicks within the same category. The piece with most clicks in category A is
assigned a score of 1. The score of the remaining items in the same category is
computed by scaling down the number of clicks by the same factor.

This results in a list where each category is represented with exactly one item
at the top of the list (each of them with score 1).

Generate a list of top 200 videos (score by normalizing popularity within each category):

    export SRG_OAUTH_TOKEN=xyz
    go run cmd/topnorm200/main.go > /tmp/topnorm-200.tsv

Sample output:

    Score	Popularity	Category	Title	Id
    1.00	376	Kids	Masha und der Bär - E churzi Reis	6b70748b-b89b-4bb9-9fec-aa6c63d291e6
    1.00	1755	Talk	Brian alias «Carlos» – Kommt er je frei?	ec3704b5-cc9b-47cf-9781-6026c61d6471
    1.00	323	Comedy	Staffelfinale mit Kilian Ziegler (Staffel 14, Folge 10)	0f4565dd-5a7f-41b9-9034-566d1e89c60c
    1.00	5608	Wissen & Ratgeber	Wie wir Behinderte behindern	31f99f6e-d61c-46e1-b785-7d4c7ed9696b
    1.00	3604	Unterhaltung	Abschied mit Tränen	e01df7e2-e401-4334-84c4-6a19ed807bd4
    1.00	1147	Serien	Vertrauen (Staffel 15, Folge 27)	1b981517-5a19-4fdc-9d68-6f010eef2ff1
    1.00	1014	Gesellschaft	Neuer TikTok-Star?: Sascha Ruefer geht mit WM-Kommentar viral	529bf8ab-7955-477d-bee3-5b9d9e752dd0
    1.00	3402	Dok & Reportagen	Wir, die Pflegefachkräfte von morgen (Staffel 1, Folge 1)	8bc3e03b-5088-4701-b860-f1726d9a0347
    1.00	2984		Gelingt die Einigung im Rentenstreit?	9e8802ad-029e-4ce9-889a-ac3c9237a964
    1.00	11112	Sport	Auf dem Lift mit «Gumpiball» Aline Danioth	bd1deb25-2ddc-43bd-b863-26a7c3b4e29a
    1.00	11276	News	Tagesschau vom 25.11.2022: Hauptausgabe	658b3c99-db7e-473b-bf17-717039cb98ac
    1.00	2724	Politik	Rentenstreit reloaded: Jetzt geht es um die zweite Säule	5f40dc56-253f-4a30-9872-996edf9b6088
    1.00	1516	Kultur	Black Friday: Kaufrausch vs. Konsumverzicht	86d70fa2-7951-4b34-a164-2bbff7d15edb
    1.00	3195	Film	Inga Lindström – Rasmus und Johanna	8af0ae11-102e-4a2d-8ffb-18e5aaed0ed4
    0.90	10016	Sport	Odermatt: «In Lake Louise gibt es nicht sehr viele Helfer»	df9dcb12-461e-4a91-918e-38704eeaaaa1
    0.85	977	Serien	Hoffnungsschimmer (Staffel 15, Folge 28)	b5b02651-d778-4596-98bf-1e69897b9517
    0.80	1412	Talk	mit Zarifa Ghafari	169e301e-d3bd-4e45-9b54-abcfe5d2049d
    ...

However, it is questionable if really every category needs to be represented
with one top-hit at the start of the ranking. Thus, let's try a mixed approach.

### Top 200 videos with clicks normalized using a mixed approach

A straightforward way to refine the score is to include more factors in the
equation. In addition to the score derived from normalizing clicks within a
category, a second factor is derived from normalizing clicks over all items.

To refine the contribution of one or the other factor to the total score, they
can be raised to the power of some exponent.

Generate a list of top 200 videos (score with mixed approach):

By default this is using the following exponents:

  * Exponent used to weight the score derived from clicks normalized over all items:
    `expGeneral=1`
  * Exponent used to weight the score derived from clicks normalized within each category:
    `expCategory=2`

Change the source code in order to adapt the exponents.

    export SRG_OAUTH_TOKEN=xyz
    go run cmd/topmixnorm200/main.go > /tmp/topmixnorm-200.tsv

Sample output:

    Score	Popularity	Category	Title	Id
    1.00	11276	News	Tagesschau vom 25.11.2022: Hauptausgabe	658b3c99-db7e-473b-bf17-717039cb98ac
    0.99	11112	Sport	Auf dem Lift mit «Gumpiball» Aline Danioth	bd1deb25-2ddc-43bd-b863-26a7c3b4e29a
    0.72	10016	Sport	Odermatt: «In Lake Louise gibt es nicht sehr viele Helfer»	df9dcb12-461e-4a91-918e-38704eeaaaa1
    0.50	5608	Wissen & Ratgeber	Wie wir Behinderte behindern	31f99f6e-d61c-46e1-b785-7d4c7ed9696b
    0.32	3604	Unterhaltung	Abschied mit Tränen	e01df7e2-e401-4334-84c4-6a19ed807bd4
    0.30	3402	Dok & Reportagen	Wir, die Pflegefachkräfte von morgen (Staffel 1, Folge 1)	8bc3e03b-5088-4701-b860-f1726d9a0347
    0.28	3195	Film	Inga Lindström – Rasmus und Johanna	8af0ae11-102e-4a2d-8ffb-18e5aaed0ed4
    0.26	2984		Gelingt die Einigung im Rentenstreit?	9e8802ad-029e-4ce9-889a-ac3c9237a964
    0.24	2724	Politik	Rentenstreit reloaded: Jetzt geht es um die zweite Säule	5f40dc56-253f-4a30-9872-996edf9b6088
    0.16	1755	Talk	Brian alias «Carlos» – Kommt er je frei?	ec3704b5-cc9b-47cf-9781-6026c61d6471
    0.13	1516	Kultur	Black Friday: Kaufrausch vs. Konsumverzicht	86d70fa2-7951-4b34-a164-2bbff7d15edb
    0.11	2317	Film	Inga Lindström – Hochzeit in Hardingsholm	123563f0-ec53-4c53-91f0-35cd19eee7ba
    0.10	1147	Serien	Vertrauen (Staffel 15, Folge 27)	1b981517-5a19-4fdc-9d68-6f010eef2ff1
    0.10	5179	Sport	Odermatt: «Hat mich selber etwas überrascht»	fd74ec5b-612e-4fb8-84ef-f0eb9a1ff9f5
    0.09	1014	Gesellschaft	Neuer TikTok-Star?: Sascha Ruefer geht mit WM-Kommentar viral	529bf8ab-7955-477d-bee3-5b9d9e752dd0
    0.08	1412	Talk	mit Zarifa Ghafari	169e301e-d3bd-4e45-9b54-abcfe5d2049d
    0.07	1191	Kultur	Mythos Matter	6702d57e-2f65-424d-8a0a-176a65b5e1f0
    0.06	977	Serien	Hoffnungsschimmer (Staffel 15, Folge 28)	b5b02651-d778-4596-98bf-1e69897b9517
    0.06	1983	Dok & Reportagen	Katar – Perlen im Sand	c00cafb8-8a9b-47cb-9053-c294bd2815cb
    0.06	4302	News	10 vor 10 vom 25.11.2022	8fb42def-a546-4c6a-b416-0d7d6e1bf959
    0.05	1644	Politik	Männer ausgeschlossen, Millionen für die Papst-Garde, Transfrau im Männergefängnis	d20389af-8ee3-43a7-be9b-fd95ba74fb38
    0.05	1943	Unterhaltung	Darum sind die Lieder von Mani Matter zeitlos | Making-of | Neu aufgelegt: Mani Matter 2022	dfb3389b-52a9-4467-92a3-87e471f7e68e
    0.05	1700		Oksana Mirinenko findet nur langsam zurück ins Leben	3851ea9f-83de-40e7-8096-459ded72f1e9
    0.04	3735	Sport	Feuz und Co. reut's, den Iltis freut's	a745ab17-f042-4db1-b340-f7a717078124

### Effects of applying the mixed approach on the ranking

Evaluating the resultis it becomes clear that a higher value of `expCategory`
leads to more content of distinct categories among the top ranked entris.

Results for `expCategory=1`:

![Screenshot of Spreadsheet with Results for `expCategory=1`](./img/ranking-effect-mixed-exp-category-1.png)

Results for `expCategory=2`:

![Screenshot of Spreadsheet with Results for `expCategory=2`](./img/ranking-effect-mixed-exp-category-2.png)
