# pocketcasts-readme
Updates README with recently listened podcasts from Pocketcasts

# Example

## Recently Listened Podcasts
<!-- PODCAST-LIST:START -->
The Knowledge Project with Shane Parrish: [#152 Tobi Lütke: Calm Progress](https://traffic.libsyn.com/secure/theknowledgeproject/152_Tobi_Lutke.mp3?dest-id=271299)

Unexplainable: [Holding on to power](https://www.podtrac.com/pts/redirect.mp3/pdst.fm/e/chtbl.com/track/524GE/traffic.megaphone.fm/VMP2092320771.mp3?updated=1667947232)

Huberman Lab: [Controlling Your Dopamine For Motivation, Focus & Satisfaction](https://www.podtrac.com/pts/redirect.mp3/pdst.fm/e/chrt.fm/track/3F7F74/traffic.megaphone.fm/SCIM6449668176.mp3?updated=1687619103)

WSJ’s The Future of Everything: [Navigating The Future of Maps](https://pdst.fm/e/chrt.fm/track/1F1B1F/traffic.megaphone.fm/WSJ6333562003.mp3?updated=1671815114)

Unexplainable: [Origins: The first living thing](https://www.podtrac.com/pts/redirect.mp3/pdst.fm/e/chtbl.com/track/524GE/traffic.megaphone.fm/VMP5597083531.mp3?updated=1678822191)


<!-- PODCAST-LIST:END -->

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
The Knowledge Project with Shane Parrish: [#152 Tobi Lütke: Calm Progress](https://traffic.libsyn.com/secure/theknowledgeproject/152_Tobi_Lutke.mp3?dest-id=271299)

Unexplainable: [Holding on to power](https://www.podtrac.com/pts/redirect.mp3/pdst.fm/e/chtbl.com/track/524GE/traffic.megaphone.fm/VMP2092320771.mp3?updated=1667947232)

Huberman Lab: [Controlling Your Dopamine For Motivation, Focus & Satisfaction](https://www.podtrac.com/pts/redirect.mp3/pdst.fm/e/chrt.fm/track/3F7F74/traffic.megaphone.fm/SCIM6449668176.mp3?updated=1687619103)

WSJ’s The Future of Everything: [Navigating The Future of Maps](https://pdst.fm/e/chrt.fm/track/1F1B1F/traffic.megaphone.fm/WSJ6333562003.mp3?updated=1671815114)

Unexplainable: [Origins: The first living thing](https://www.podtrac.com/pts/redirect.mp3/pdst.fm/e/chtbl.com/track/524GE/traffic.megaphone.fm/VMP5597083531.mp3?updated=1678822191)


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
