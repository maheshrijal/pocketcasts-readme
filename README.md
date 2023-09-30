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
The Knowledge Project with Shane Parrish: [#174 Rey Flemings: A Different Definition of Success](https://traffic.libsyn.com/secure/theknowledgeproject/174_Rey_Flemings.mp3?dest-id=271299)

Lex Fridman Podcast: [#396 – James Sexton: Divorce Lawyer on Marriage, Relationships, Sex, Lies & Love](https://media.blubrry.com/takeituneasy/content.blubrry.com/takeituneasy/lex_ai_james_sexton.mp3)

The Changelog: Software Development, Open Source: [Efficient Linux at the CLI](https://op3.dev/e/https://cdn.changelog.com/uploads/podcast/547/the-changelog-547.mp3)

3 Things: [The Catch Up: 29 September](https://dts.podtrac.com/redirect.mp3/api.spreaker.com/download/episode/56992111/catch_up_2023_29th_september_v1.mp3)

The Changelog: Software Development, Open Source: [The missing sync layer for modern apps](https://op3.dev/e/https://cdn.changelog.com/uploads/news/63/changelog-news-63.mp3)


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
