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
Capitalmind Podcast: [How to invest a lumpsum amount?](https://traffic.libsyn.com/secure/capitalmind/CM_Podcast_EP_68_0816.mp3?dest-id=1084508)

The Knowledge Project with Shane Parrish: [#127 Best of 2021: Conversations of the Year](https://traffic.libsyn.com/secure/theknowledgeproject/Episode_127_-_Best_of_2021.mp3?dest-id=271299)

Finshots Daily: [Why do creditors misuse the Insolvency Code?](https://anchor.fm/s/37a76020/podcast/play/76319379/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2023-8-24%2Fb30d0670-91be-613a-61de-9ca4e710c100.mp3)

Lex Fridman Podcast: [#396 – James Sexton: Divorce Lawyer on Marriage, Relationships, Sex, Lies & Love](https://media.blubrry.com/takeituneasy/content.blubrry.com/takeituneasy/lex_ai_james_sexton.mp3)

Finshots Daily: [India wants a bigger bite of the MICE tourism pie](https://anchor.fm/s/37a76020/podcast/play/76219934/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2023-8-22%2Fd804103f-879b-5cc8-7f88-251ed61289de.mp3)


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
