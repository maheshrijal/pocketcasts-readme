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
Huberman Lab: [Chris Voss: How to Succeed at Hard Conversations](https://www.podtrac.com/pts/redirect.mp3/pdst.fm/e/chrt.fm/track/3F7F74/traffic.megaphone.fm/SCIM5142864264.mp3?updated=1696225071)

Finshots Daily: [Can a tax solve Bengaluru’s traffic woes?](https://anchor.fm/s/37a76020/podcast/play/76939995/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2023-9-8%2Fa5faa02d-35e3-8ac0-126c-bc4fef90c00d.mp3)

The Knowledge Project with Shane Parrish: [#72 Neil Pasricha: Happy Habits](https://traffic.libsyn.com/secure/theknowledgeproject/KP20Ep2007220Neil20Pasricha20Audio20Master20Rev1.mp3?dest-id=271299)

Finshots Daily: [The tussle between bankrupt airlines and lessors is over. Almost.](https://anchor.fm/s/37a76020/podcast/play/76831914/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2023-9-5%2F349918729-44100-2-83cd7a7bdbed6.m4a)

Lex Fridman Podcast: [#210 – Matt Walker: Sleep](https://media.blubrry.com/takeituneasy/content.blubrry.com/takeituneasy/lex_ai_matt_walker.mp3)


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
