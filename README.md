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
The Knowledge Project with Shane Parrish: [#72 Neil Pasricha: Happy Habits](https://traffic.libsyn.com/secure/theknowledgeproject/KP20Ep2007220Neil20Pasricha20Audio20Master20Rev1.mp3?dest-id=271299)

Lex Fridman Podcast: [#210 – Matt Walker: Sleep](https://media.blubrry.com/takeituneasy/content.blubrry.com/takeituneasy/lex_ai_matt_walker.mp3)

Murder in Apartment 12: [Introducing: Murder in Apartment 12](https://chrt.fm/track/6D589D/dts.podtrac.com/redirect.mp3/nbcnews.simplecastaudio.com/079d7701-3b3e-4e63-8bad-98cf272a9638/episodes/9e48db29-30bf-410a-a780-8f62a7797c19/audio/128/default.mp3?aid=rss_feed&awCollectionId=079d7701-3b3e-4e63-8bad-98cf272a9638&awEpisodeId=9e48db29-30bf-410a-a780-8f62a7797c19&feed=RPWEjhKq)

Three Minute Masters: [Robin Adams on Game Changing Automation](https://sphinx.acast.com/p/open/s/6411efc6467ff50011b81aa3/e/642ab06013df9d0011c5a402/media.mp3)

The Morgan Housel Podcast: [Death, Taxes, and a Few Other Things](https://www.buzzsprout.com/2144602/13731114-death-taxes-and-a-few-other-things.mp3)


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
