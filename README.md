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
Capitalmind Podcast: [Should you invest in a PMS?](https://traffic.libsyn.com/secure/capitalmind/CM_EP_69_1010_Final.mp3?dest-id=1084508)

Huberman Lab: [Chris Voss: How to Succeed at Hard Conversations](https://www.podtrac.com/pts/redirect.mp3/pdst.fm/e/chrt.fm/track/3F7F74/traffic.megaphone.fm/SCIM5142864264.mp3?updated=1696225071)

Finshots Daily: [Adani versus Financial Times?](https://anchor.fm/s/37a76020/podcast/play/77262433/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2023-9-16%2F9f1926da-2143-8b43-9dee-3a952e73e02c.mp3)

The Changelog: Software Development, Open Source: [InfluxDB drops Go for Rust but gokrazy is really cool](https://op3.dev/e/https://cdn.changelog.com/uploads/news/64/changelog-news-64.mp3)

The Changelog: Software Development, Open Source: [RTO vs WFH & the case for strong static typing](https://op3.dev/e/https://cdn.changelog.com/uploads/news/65/changelog-news-65.mp3)


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
