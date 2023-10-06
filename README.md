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
Darknet Diaries: [The Mimics of Punjab](https://www.podtrac.com/pts/redirect.mp3/pdst.fm/e/chrt.fm/track/G481GD/traffic.megaphone.fm/ADV7639639710.mp3?updated=1696273602)

Finshots Daily: [Can Kashmir end the English monopoly?](https://anchor.fm/s/37a76020/podcast/play/76732417/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2023-9-3%2F13920c3d-4d1c-13d4-ad31-ea85997268b4.mp3)

Finshots Daily: [An Apple Search Engine?](https://anchor.fm/s/37a76020/podcast/play/76783803/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2023-9-4%2F07530a1f-422e-4192-aa13-696ebc272e93.mp3)

Lex Fridman Podcast: [#264 – Tim Urban: Elon Musk, Neuralink, AI, Aliens, and the Future of Humanity](https://media.blubrry.com/takeituneasy/content.blubrry.com/takeituneasy/lex_ai_tim_urban.mp3)

Finshots Daily: [A story about MS Swaminathan and India’s Green Revolution](https://anchor.fm/s/37a76020/podcast/play/76694565/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2023-9-2%2F780f3ed1-d6b6-454a-0035-65c168ec9757.mp3)


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
