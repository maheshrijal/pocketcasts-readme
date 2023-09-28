# pocketcasts-readme
Updates README with recently listened podcasts from Pocketcasts

# Setup

## Prep Work

1. You need to save the pocketcasts Email & Password in the repository secrets. You can find that in the Settings of your Repository. Be sure to save those as the following.

    - `POCKETCASTS_EMAIL = <your  Pocketcasts Email>`
    - `POCKETCASTS_PASSWORD = <your  Pocketcasts Password>`

# Update your README

Add a comment to your README.md like this:

```markdown
# # Recently Listened Podcasts

<!-- PODCAST-LIST:START -->
Lex Fridman Podcast: [#396 – James Sexton: Divorce Lawyer on Marriage, Relationships, Sex, Lies & Love](https://media.blubrry.com/takeituneasy/content.blubrry.com/takeituneasy/lex_ai_james_sexton.mp3)

Optimal Finance Daily: Money Management & Financial Independence: [Motivation and Money by Honey Smith of Get Rich Slowly on Self-Awareness and Finances](https://chrt.fm/track/G52973/traffic.megaphone.fm/OLD1859299240.mp3?updated=1695055694)

IRL: Online Life is Real Life: [We’re Back! IRL Season 7: People Over Profit](https://cdn.simplecast.com/audio/9b52b824-909f-4be5-aaf0-10f9e93c7818/episodes/41c243fa-78eb-4580-a33c-826e74ff3e40/audio/a6355dbb-4085-4c91-9b21-641facd43d61/default_tc.mp3?aid=rss_feed&feed=lP7owBq8)

IRL: Online Life is Real Life: [The AI Medicine Cabinet](https://cdn.simplecast.com/audio/9b52b824-909f-4be5-aaf0-10f9e93c7818/episodes/cab8b72a-383d-4843-a0bc-7172729eb7ce/audio/3df09b0a-7c8e-40b2-b2a9-bac320d799e3/default_tc.mp3?aid=rss_feed&feed=lP7owBq8)

Planet Money: ["Based on a true story"](https://play.podtrac.com/npr-510289/traffic.megaphone.fm/NPR7608670799.mp3?d=1735&size=27771342&e=1197954107&t=podcast&p=510289)


<!-- PODCAST-LIST:END -->
```

# Repository Workflow

Please follow the steps below:

1. Go to your `<username>/<username>/actions`, hit New workflow, set up a workflow yourself, delete all the default content github made for you.
2. Copy the below code and paste it to your new workflow


```
name: Update pods
on:
    workflow_dispatch:
    schedule:
    # Runs every 8 hours
    - cron: "0 */8 * * *"

permissions:
    contents: write

jobs:
    build:
        runs-on: ubuntu-latest

        steps:
            - name: Checkout Code
              uses: actions/checkout@v3
            - name: Update README
              uses: maheshrjl/pocketcasts-readme@main
              with:
                  POCKETCASTS_EMAIL: ${{ secrets.POCKETCASTS_EMAIL }}
                  POCKETCASTS_PASSWORD: ${{ secrets.POCKETCASTS_PASSWORD }}
            - name: Push Changes
              uses: stefanzweifel/git-auto-commit-action@v4
```


## License

[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)

To the extent possible under law, [Mahesh Rijal](https://maheshrjl.com/) has waived all copyright and related or neighboring rights to this work.

_Inspired by [abhisheknaiidu/todoist-readme](https://github.com/abhisheknaiidu/todoist-readme)_
