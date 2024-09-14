<div align="center">

<img src="./images/fabric-logo-gif.gif" alt="fabriclogo" width="400" height="400"/>

# `fabric`

![Static Badge](https://img.shields.io/badge/mission-human_flourishing_via_AI_augmentation-purple)
<br />
![GitHub top language](https://img.shields.io/github/languages/top/danielmiessler/fabric)
![GitHub last commit](https://img.shields.io/github/last-commit/danielmiessler/fabric)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

<p class="align center">
<h4><code>fabric</code> is an open-source framework for augmenting humans using AI.</h4>
</p>

[What and Why](#whatandwhy) •
[Philosophy](#philosophy) •
[Installation](#Installation) •
[Usage](#Usage) •
[Examples](#examples) •
[Just Use the Patterns](#just-use-the-patterns) •
[Custom Patterns](#custom-patterns) •
[Helper Apps](#helper-apps) •
<<<<<<< HEAD


<img src="./images/fabric-logo-gif.gif" alt="fabriclogo" width="400" height="400"/>

# `fabric`

<h4><code>fabric</code> is an open-source framework for augmenting humans with AI.</h4>

[Description](#description) •
[Naming](#naming) •
[Structure](#structure) •
[Functionality](#functionality) •
[Installation](#installation) •
[Usage](#usage) •
[Examples](#examples) •
=======
>>>>>>> 9b8871f25b17dfeb7338ebaf39078bb0929b0fb9
[Meta](#meta)

</div>

## Navigation

- [What and Why](#what-and-why)
- [Philosophy](#philosophy)
  - [Breaking problems into components](#breaking-problems-into-components)
  - [Too many prompts](#too-many-prompts)
  - [The Fabric approach to prompting](#our-approach-to-prompting)
- [Installation](#Installation)
  - [Migration](#Migration)
  - [Upgrading](#Upgrading)
- [Usage](#Usage)
- [Examples](#examples)
  - [Just use the Patterns](#just-use-the-patterns)
- [Custom Patterns](#custom-patterns)
- [Helper Apps](#helper-apps)
- [Meta](#meta)
  - [Primary contributors](#primary-contributors)

<br />

> [!NOTE] 
August 20, 2024 — We have migrated to Go, and the transition has been pretty smooth! The biggest thing to know is that **the previous installation instructions in the various Fabric videos out there will no longer work** because they were for the legacy (Python) version. Check the new [install instructions](#Installation) below.
>
>
> **The following command line options were changed during the migration to Go:**
> * You now need to use the -c option instead of -C to copy the result to the clipboard.
> * You now need to use the -s option instead of -S to stream results in realtime.
> * The following command line options have been removed --agents (-a), --gui, --clearsession, --remoteOllamaServer, and --sessionlog options 
> * You can now use --Setup (-S) to configure an Ollama server.
> * **Please be patient while our developers rewrite the gui in go**

## Intro videos

Keep in mind that many of these were recorded when Fabric was Python-based, so remember to use the current [install instructions](#Installation) below.

* [Network Chuck](https://www.youtube.com/watch?v=UbDyjIIGaxQ)
* [David Bombal](https://www.youtube.com/watch?v=vF-MQmVxnCs)
* [My Own Intro to the Tool](https://www.youtube.com/watch?v=wPEyyigh10g)
* [More Fabric YouTube Videos](https://www.youtube.com/results?search_query=fabric+ai)

## What and why

Since the start of 2023 and GenAI we've seen a massive number of AI applications for accomplishing tasks. It's powerful, but _it's not easy to integrate this functionality into our lives._

<div align="center">
<h4>In other words, AI doesn't have a capabilities problem—it has an <em>integration</em> problem.</h4>
</div>

Fabric was created to address this by enabling everyone to granularly apply AI to everyday challenges.

## Philosophy

> AI isn't a thing; it's a _magnifier_ of a thing. And that thing is **human creativity**.

We believe the purpose of technology is to help humans flourish, so when we talk about AI we start with the **human** problems we want to solve.

### Breaking problems into components

Our approach is to break problems into individual pieces (see below) and then apply AI to them one at a time. See below for some examples.

<img width="2078" alt="augmented_challenges" src="https://github.com/danielmiessler/fabric/assets/50654/31997394-85a9-40c2-879b-b347e4701f06">

### Too many prompts

Prompts are good for this, but the biggest challenge I faced in 2023——which still exists today—is **the sheer number of AI prompts out there**. We all have prompts that are useful, but it's hard to discover new ones, know if they are good or not, _and manage different versions of the ones we like_.

One of <code>fabric</code>'s primary features is helping people collect and integrate prompts, which we call _Patterns_, into various parts of their lives.

Fabric has Patterns for all sorts of life and work activities, including:

- Extracting the most interesting parts of YouTube videos and podcasts
- Writing an essay in your own voice with just an idea as an input
- Summarizing opaque academic papers
- Creating perfectly matched AI art prompts for a piece of writing
- Rating the quality of content to see if you want to read/watch the whole thing
- Getting summaries of long, boring content
- Explaining code to you
- Turning bad documentation into usable documentation
- Creating social media posts from any content input
- And a million more…

## Installation

To install Fabric, [make sure Go is installed](https://go.dev/doc/install), and then run the following command.

```bash
# Install Fabric directly from the repo
go install github.com/danielmiessler/fabric@latest
```

### Environment Variables

You may need to set some environment variables in your `~/.bashrc` on linux or `~/.zshrc` file on mac to be able to run the `fabric` command. Here is an example of what you can add:

For Intel based macs or linux
```bash
# Golang environment variables
export GOROOT=/usr/local/go
export GOPATH=$HOME/go

# Update PATH to include GOPATH and GOROOT binaries
export PATH=$GOPATH/bin:$GOROOT/bin:$HOME/.local/bin:$PATH
```

for Apple Silicon based macs
```bash
# Golang environment variables
export GOROOT=/opt/homebrew/bin/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$HOME/.local/bin:$PATH:
```

### Setup
Now run the following command
```bash
# Run the setup to set up your directories and keys
fabric --setup
```
If everything works you are good to go.


### Migration

If you have the Legacy (Python) version installed and want to migrate to the Go version, here's how you do it. It's basically two steps: 1) uninstall the Python version, and 2) install the Go version.

```bash
# Uninstall Legacy Fabric
pipx uninstall fabric

# Clear any old Fabric aliases
(check your .bashrc, .zshrc, etc.)
# Install the Go version
go install github.com/danielmiessler/fabric@latest
# Run setup for the new version. Important because things have changed
fabric --setup
```

Then [set your environmental variables](#environmental-variables) as shown above.

### Upgrading

The great thing about Go is that it's super easy to upgrade. Just run the same command you used to install it in the first place and you'll always get the latest version.
```bash
go install github.com/danielmiessler/fabric@latest
```

## Usage
Once you have it all set up, here's how to use it.

```bash
fabric -h
```

```bash
usage: fabric -h
Usage:
  fabric [OPTIONS]

Application Options:
  -p, --pattern=                    Choose a pattern
  -v, --variable=                   Values for pattern variables, e.g. -v=$name:John -v=$age:30
  -C, --context=                    Choose a context
      --session=                    Choose a session
  -S, --setup                       Run setup
      --setup-skip-update-patterns  Skip update patterns at setup
  -t, --temperature=                Set temperature (default: 0.7)
  -T, --topp=                       Set top P (default: 0.9)
  -s, --stream                      Stream
  -P, --presencepenalty=            Set presence penalty (default: 0.0)
  -F, --frequencypenalty=           Set frequency penalty (default: 0.0)
  -l, --listpatterns                List all patterns
  -L, --listmodels                  List all available models
  -x, --listcontexts                List all contexts
  -X, --listsessions                List all sessions
  -U, --updatepatterns              Update patterns
  -c, --copy                        Copy to clipboard
  -m, --model=                      Choose model
  -o, --output=                     Output to file
  -n, --latest=                     Number of latest patterns to list (default: 0)
  -d, --changeDefaultModel          Change default pattern
  -y, --youtube=                    YouTube video url to grab transcript, comments from it and send to chat
      --transcript                  Grab transcript from YouTube video and send to chat
      --comments                    Grab comments from YouTube video and send to chat
      --dry-run                     Show what would be sent to the model without actually sending it

Help Options:
  -h, --help                        Show this help message

```

## Our approach to prompting

Fabric _Patterns_ are different than most prompts you'll see.

- **First, we use `Markdown` to help ensure maximum readability and editability**. This not only helps the creator make a good one, but also anyone who wants to deeply understand what it does. _Importantly, this also includes the AI you're sending it to!_

Here's an example of a Fabric Pattern.

```bash
https://github.com/danielmiessler/fabric/blob/main/patterns/extract_wisdom/system.md
```

<img width="1461" alt="pattern-example" src="https://github.com/danielmiessler/fabric/assets/50654/b910c551-9263-405f-9735-71ca69bbab6d">

- **Next, we are extremely clear in our instructions**, and we use the Markdown structure to emphasize what we want the AI to do, and in what order.

- **And finally, we tend to use the System section of the prompt almost exclusively**. In over a year of being heads-down with this stuff, we've just seen more efficacy from doing that. If that changes, or we're shown data that says otherwise, we will adjust.

## Examples

Now let's look at some things you can do with Fabric.

1. Run the `summarize` Pattern based on input from `stdin`. In this case, the body of an article.

```bash
pbpaste | fabric --pattern summarize
```

2. Run the `analyze_claims` Pattern with the `--stream` option to get immediate and streaming results.

```bash
pbpaste | fabric --stream --pattern analyze_claims
```

3. Run the `extract_wisdom` Pattern with the `--stream` option to get immediate and streaming results from any Youtube video (much like in the original introduction video).

```bash
yt --transcript https://youtube.com/watch?v=uXs-zPc63kM | fabric --stream --pattern extract_wisdom
```

4. Create patterns- you must create a .md file with the pattern and save it to ~/.config/fabric/patterns/[yourpatternname].

## Just use the Patterns

<img width="1173" alt="fabric-patterns-screenshot" src="https://github.com/danielmiessler/fabric/assets/50654/9186a044-652b-4673-89f7-71cf066f32d8">

<br />
<br />

If you're not looking to do anything fancy, and you just want a lot of great prompts, you can navigate to the [`/patterns`](https://github.com/danielmiessler/fabric/tree/main/patterns) directory and start exploring!

We hope that if you used nothing else from Fabric, the Patterns by themselves will make the project useful.

You can use any of the Patterns you see there in any AI application that you have, whether that's ChatGPT or some other app or website. Our plan and prediction is that people will soon be sharing many more than those we've published, and they will be way better than ours.

The wisdom of crowds for the win.

## Custom Patterns

You may want to use Fabric to create your own custom Patterns—but not share them with others. No problem!

Just make a directory in `~/.config/custompatterns/` (or wherever) and put your `.md` files in there. 

When you're ready to use them, copy them into:

```
~/.config/fabric/patterns/
```

You can then use them like any other Patterns, but they won't be public unless you explicitly submit them as Pull Requests to the Fabric project. So don't worry—they're private to you.


This feature works with all openai and ollama models but does NOT work with claude. You can specify your model with the -m flag

## Helper Apps

Fabric also makes use of some core helper apps (tools) to make it easier to integrate with your various workflows. Here are some examples:

`yt` is a helper command that extracts the transcript from a YouTube video. You can use it like this:
```bash
yt https://www.youtube.com/watch?v=lQVcbY52_gY
```

This will return the transcript from the video, which you can then pipe into Fabric like this:
```bash
yt https://www.youtube.com/watch?v=lQVcbY52_gY | fabric --pattern extract_wisdom
```

### `yt` Installation

To install `yt`, install it the same way as you install Fabric, just with a different repo name.

```bash
go install github.com/danielmiessler/yt@latest
```

Be sure to add your `YOUTUBE_API_KEY` to `~/.config/fabric/.env`.

### `to_pdf`

`to_pdf` is a helper command that converts LaTeX files to PDF format. You can use it like this:

```bash
to_pdf input.tex
```

This will create a PDF file from the input LaTeX file in the same directory.

You can also use it with stdin which works perfectly with the `write_latex` pattern:

```bash
echo "ai security primer" | fabric --pattern write_latex | to_pdf
```

This will create a PDF file named `output.pdf` in the current directory.

### `to_pdf` Installation

To install `to_pdf`, install it the same way as you install Fabric, just with a different repo name.

```bash
go install github.com/danielmiessler/fabric/to_pdf/to_pdf@latest
```

Make sure you have a LaTeX distribution (like TeX Live or MiKTeX) installed on your system, as `to_pdf` requires `pdflatex` to be available in your system's PATH.

## Meta

> [!NOTE]
> Special thanks to the following people for their inspiration and contributions!

- _Jonathan Dunn_ for being the absolute MVP dev on the project, including spearheading the new Go version, as well as the GUI! All this while also being a full-time medical doctor!
- _Caleb Sima_ for pushing me over the edge of whether to make this a public project or not.
- _Eugen Eisler_ and _Frederick Ros_ for their invaluable contributions to the Go version
- _Joel Parish_ for super useful input on the project's Github directory structure..
- _Joseph Thacker_ for the idea of a `-c` context flag that adds pre-created context in the `./config/fabric/` directory to all Pattern queries.
- _Jason Haddix_ for the idea of a stitch (chained Pattern) to filter content using a local model before sending on to a cloud model, i.e., cleaning customer data using `llama2` before sending on to `gpt-4` for analysis.
- _Andre Guerra_ for assisting with numerous components to make things simpler and more maintainable.

### Primary contributors

<a href="https://github.com/danielmiessler"><img src="https://avatars.githubusercontent.com/u/50654?v=4" title="Daniel Miessler" width="50" height="50"></a>
<a href="https://github.com/xssdoctor"><img src="https://avatars.githubusercontent.com/u/9218431?v=4" title="Jonathan Dunn" width="50" height="50"></a>
<a href="https://github.com/sbehrens"><img src="https://avatars.githubusercontent.com/u/688589?v=4" title="Scott Behrens" width="50" height="50"></a>
<a href="https://github.com/agu3rra"><img src="https://avatars.githubusercontent.com/u/10410523?v=4" title="Andre Guerra" width="50" height="50"></a>

`fabric` was created by <a href="https://danielmiessler.com/subscribe" target="_blank">Daniel Miessler</a> in January of 2024.
<br /><br />
<a href="https://twitter.com/intent/user?screen_name=danielmiessler">![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/danielmiessler)</a>
<<<<<<< HEAD
=======
<br />

## ⚠️ This project is currently pre-launch—with code, functionality, and documentation still being added. But any published patterns can be used in the meantime.⚠️

## Description

Since the end of 2023 we've seen a massive number of different AI applications for accomplishing tasks. The problem is that it's not easy to integrate them with our lives.

Perhaps the greatest example is the number of AI prompts out there. We all have prompts that are useful, but it's hard to manage them, discover new ones, and track the different versions of the ones we like.

One of <code>fabric</code>'s main features is helping people collect and integrate modular AI functionality, which we call _Patterns_, into various parts of their lives.

There are patterns for all sorts of life and work activities, including:

- Extracting the most interesting parts of YouTube videos and podcasts
- Summarizing opaque academic papers
- Creating perfectly matched AI art prompts for a piece of writing
- Rating the quality of content to see if you want to read/watch the whole thing
- Getting summaries of long, boring content
- Write a full essay in a particular voice, given an idea as input
- Create social media posts from any content input
- And a million more…

## Naming

`fabric` is themed off of, um…_fabric_. So, think blankets, quilts, patterns, etc. Here's the structure:

- The project itself is called **Fabric**, and it's the parent concept.
- Individual AI modules (think prompts) are called **Patterns**.
- Chaining together _Patterns_ to create advanced functionality is called a **Stitch**.
- The optional server-side functionality of `fabric` is called the **Mill**.
- The optional client-side scripts within `fabric` are called **Looms**.

## Functionality

`fabric`'s main function is to make **Patterns** available to everyone in an open ecosystem, i.e., to allow people to share and fork prompts in a transparent, scalable, and dependable way.

But it also includes two other components that make it possible for AI enthusiasts and developers to _build your own Personal AI Ecosystem_.

_In other words you can have your own server, with your own copy of `fabric`, running your own specific combination of **Patterns** for your specific use cases._

Here are the three `fabric` ecosystem pieces, and how they work together.

- The **Mill** is the (optional) server that makes **Patterns** available.
- **Patterns** are the actual AI use cases.
- **Looms** are the modular, client-side apps that call a specific **Pattern\* hosted by a **Mill\*\*.

## Usage

One key feature of `fabric` and its Markdown-based format is the ability to reference (and edit) individual [patterns](https://github.com/danielmiessler/fabric/tree/main#naming) directly—on their own—without surrounding code.

As an example, heres how to call _the direct location_ of the **system** prompt for the `extractwisdom` pattern.

```sh
https://github.com/danielmiessler/fabric/blob/main/extract-wisdom/dmiessler/extract-wisdom-1.0.0/system.md
```

This means you can cleanly, and directly reference any pattern within `fabric`, for use in a web-based AI app, your own programming, or wherever!

Even better, you can also have your [Mill](https://github.com/danielmiessler/fabric/tree/main#naming) functionality directly call **system** and **user** prompts from `fabric`, meaning you can have your personal AI ecosystem automatically kept up to date with the latest version of your favorite [Patterns](https://github.com/danielmiessler/fabric/tree/main#naming).

## Examples

Here's an abridged ouptut example from the <a href="https://github.com/danielmiessler/fabric/tree/main/extract-wisdom/dmiessler/extract-wisdom-1.0.0">`extractwisdom`</a> pattern (limited to only 10 items per section).

```markdown
## SUMMARY:

The content features a conversation between two individuals discussing various topics, including the decline of Western culture, the importance of beauty and subtlety in life, the impact of technology and AI, the resonance of Rilke's poetry, the value of deep reading and revisiting texts, the captivating nature of Ayn Rand's writing, the role of philosophy in understanding the world, and the influence of drugs on society. They also touch upon creativity, attention spans, and the importance of introspection.

## IDEAS:

1. Western culture is perceived to be declining due to a loss of values and an embrace of mediocrity.
2. Mass media and technology have contributed to shorter attention spans and a need for constant stimulation.
3. Rilke's poetry resonates due to its focus on beauty and ecstasy in everyday objects.
4. Subtlety is often overlooked in modern society due to sensory overload.
5. The role of technology in shaping music and performance art is significant.
6. Reading habits have shifted from deep, repetitive reading to consuming large quantities of new material.
7. Revisiting influential books as one ages can lead to new insights based on accumulated wisdom and experiences.
8. Fiction can vividly illustrate philosophical concepts through characters and narratives.
9. Many influential thinkers have backgrounds in philosophy, highlighting its importance in shaping reasoning skills.
10. Philosophy is seen as a bridge between theology and science, asking questions that both fields seek to answer.

## QUOTES:

1. "You can't necessarily think yourself into the answers. You have to create space for the answers to come to you."
2. "The West is dying and we are killing her."
3. "The American Dream has been replaced by mass packaged mediocrity porn, encouraging us to revel like happy pigs in our own meekness."
4. "There's just not that many people who have the courage to reach beyond consensus and go explore new ideas."
5. "I'll start watching Netflix when I've read the whole of human history."
6. "Rilke saw beauty in everything... He sees it's in one little thing, a representation of all things that are beautiful."
7. "Vanilla is a very subtle flavor... it speaks to sort of the sensory overload of the modern age."
8. "When you memorize chapters [of the Bible], it takes a few months, but you really understand how things are structured."
9. "As you get older, if there's books that moved you when you were younger, it's worth going back and rereading them."
10. "She [Ayn Rand] took complicated philosophy and embodied it in a way that anybody could resonate with."

## HABITS:

1. Avoiding mainstream media consumption for deeper engagement with historical texts and personal research.
2. Regularly revisiting influential books from youth to gain new insights with age.
3. Engaging in deep reading practices rather than skimming or speed-reading material.
4. Memorizing entire chapters or passages from significant texts for better understanding.
5. Disengaging from social media and fast-paced news cycles for more focused thought processes.
6. Walking long distances as a form of meditation and reflection.
7. Creating space for thoughts to solidify through introspection and stillness.
8. Embracing emotions such as grief or anger fully rather than suppressing them.
9. Seeking out varied experiences across different careers and lifestyles.
10. Prioritizing curiosity-driven research without specific goals or constraints.

## FACTS:

1. The West is perceived as declining due to cultural shifts away from traditional values.
2. Attention spans have shortened due to technological advancements and media consumption habits.
3. Rilke's poetry emphasizes finding beauty in everyday objects through detailed observation.
4. Modern society often overlooks subtlety due to sensory overload from various stimuli.
5. Reading habits have evolved from deep engagement with texts to consuming large quantities quickly.
6. Revisiting influential books can lead to new insights based on accumulated life experiences.
7. Fiction can effectively illustrate philosophical concepts through character development and narrative arcs.
8. Philosophy plays a significant role in shaping reasoning skills and understanding complex ideas.
9. Creativity may be stifled by cultural nihilism and protectionist attitudes within society.
10. Short-term thinking undermines efforts to create lasting works of beauty or significance.

## REFERENCES:

1. Rainer Maria Rilke's poetry
2. Netflix
3. Underworld concert
4. Katy Perry's theatrical performances
5. Taylor Swift's performances
6. Bible study
7. Atlas Shrugged by Ayn Rand
8. Robert Pirsig's writings
9. Bertrand Russell's definition of philosophy
10. Nietzsche's walks
```

## Meta

`fabric` was created by <a href="https://danielmiessler.com/" target="_blank">Daniel Miessler</a>.

Special thanks to the following people for inspiration and contributions.

- **Caleb Sima** for pushing me over the edge of whether to make this a public project or not.
- **Joel Parish** for super useful input on the project's Github directory structure.

```

```

=======
>>>>>>> 9b8871f25b17dfeb7338ebaf39078bb0929b0fb9
