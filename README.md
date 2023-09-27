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
Planet Money: ["Based on a true story"](https://play.podtrac.com/npr-510289/traffic.megaphone.fm/NPR7608670799.mp3?d=1735&size=27771342&e=1197954107&t=podcast&p=510289)

The Knowledge Project with Shane Parrish: [#127 Best of 2021: Conversations of the Year](https://traffic.libsyn.com/secure/theknowledgeproject/Episode_127_-_Best_of_2021.mp3?dest-id=271299)

3 Things: [The Catch Up: 25 September](https://dts.podtrac.com/redirect.mp3/api.spreaker.com/download/episode/56924236/catch_up_2023_25th_september_v1.mp3)

Finshots Daily: [The impact of long notice periods](https://anchor.fm/s/37a76020/podcast/play/76368895/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2023-8-25%2F6f108935-585e-edc7-f1fa-bd772b3d3b99.mp3)

Capitalmind Podcast: [How to invest a lumpsum amount?](https://traffic.libsyn.com/secure/capitalmind/CM_Podcast_EP_68_0816.mp3?dest-id=1084508)


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
