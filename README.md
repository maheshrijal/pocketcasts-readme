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

Finshots Daily: [Why Claudia Goldin won the Nobel Prize in Economics](https://anchor.fm/s/37a76020/podcast/play/77036575/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2023-9-10%2F5b21d286-c96a-0ce9-0056-4dfa8773ff2a.mp3)

Finshots Daily: [Is Airbnb broken?](https://anchor.fm/s/37a76020/podcast/play/76989548/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2023-9-9%2F347cf6e5-9037-5764-3bef-898d56ddf988.mp3)

The Knowledge Project with Shane Parrish: [#167 Dr. Gina Poe: The Science of Better Sleep](https://traffic.libsyn.com/secure/theknowledgeproject/167_Gina_Poe.mp3?dest-id=271299)

The Knowledge Project with Shane Parrish: [#174 Rey Flemings: A Different Definition of Success](https://traffic.libsyn.com/secure/theknowledgeproject/174_Rey_Flemings.mp3?dest-id=271299)


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
