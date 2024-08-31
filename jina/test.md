[1] Title: GitHub - danielmiessler/fabric: fabric is an open-source framework for augmenting humans using AI. It provides a modular framework for solving specific problems using a crowdsourced set of AI prompts that can be used anywhere.
[1] URL Source: https://github.com/danielmiessler/fabric
[1] Description: fabric is <strong>an open-source framework for augmenting humans using AI</strong>. It provides a modular framework for solving specific problems using a crowdsourced set of AI prompts that can be used anywhere. - danielmiessler/fabric
[1] Markdown Content:
[![Image 1: fabriclogo](https://github.com/danielmiessler/fabric/raw/main/images/fabric-logo-gif.gif)](https://github.com/danielmiessler/fabric/blob/main/images/fabric-logo-gif.gif)

`fabric`
--------

[](https://github.com/danielmiessler/fabric#fabric)

[![Image 2: Static Badge](https://camo.githubusercontent.com/f022dc8d1303fe26ee78cd88e60920ef2d1baf96c629d782e8117faa8899e319/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6d697373696f6e2d68756d616e5f666c6f7572697368696e675f7669615f41495f6175676d656e746174696f6e2d707572706c65)](https://camo.githubusercontent.com/f022dc8d1303fe26ee78cd88e60920ef2d1baf96c629d782e8117faa8899e319/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6d697373696f6e2d68756d616e5f666c6f7572697368696e675f7669615f41495f6175676d656e746174696f6e2d707572706c65)  
[![Image 3: GitHub top language](https://camo.githubusercontent.com/5ebdbd74bb4ac3d78b0970aad6fde0e7ab273c96e28180744b16fc28a1ef109c/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f6c616e6775616765732f746f702f64616e69656c6d696573736c65722f666162726963)](https://camo.githubusercontent.com/5ebdbd74bb4ac3d78b0970aad6fde0e7ab273c96e28180744b16fc28a1ef109c/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f6c616e6775616765732f746f702f64616e69656c6d696573736c65722f666162726963) [![Image 4: GitHub last commit](https://camo.githubusercontent.com/a5f1f93e3ce4592ed5c1ceec7e552729cc451a843598c6f76e85c4610863fc4d/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f6c6173742d636f6d6d69742f64616e69656c6d696573736c65722f666162726963)](https://camo.githubusercontent.com/a5f1f93e3ce4592ed5c1ceec7e552729cc451a843598c6f76e85c4610863fc4d/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f6c6173742d636f6d6d69742f64616e69656c6d696573736c65722f666162726963) [![Image 5: License: MIT](https://camo.githubusercontent.com/28f4d479bf0a9b033b3a3b95ab2adc343da448a025b01aefdc0fbc7f0e169eb8/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4c6963656e73652d4d49542d677265656e2e737667)](https://opensource.org/licenses/MIT)

#### `fabric` is an open-source framework for augmenting humans using AI.

[](https://github.com/danielmiessler/fabric#fabric-is-an-open-source-framework-for-augmenting-humans-using-ai)

[What and Why](https://github.com/danielmiessler/fabric#whatandwhy) • [Philosophy](https://github.com/danielmiessler/fabric#philosophy) • [Installation](https://github.com/danielmiessler/fabric#Installation) • [Usage](https://github.com/danielmiessler/fabric#Usage) • [Examples](https://github.com/danielmiessler/fabric#examples) • [Just Use the Patterns](https://github.com/danielmiessler/fabric#just-use-the-patterns) • [Custom Patterns](https://github.com/danielmiessler/fabric#custom-patterns) • [Helper Apps](https://github.com/danielmiessler/fabric#helper-apps) • [Meta](https://github.com/danielmiessler/fabric#meta)

Navigation
----------

[](https://github.com/danielmiessler/fabric#navigation)

*   [What and Why](https://github.com/danielmiessler/fabric#what-and-why)
*   [Philosophy](https://github.com/danielmiessler/fabric#philosophy)
    *   [Breaking problems into components](https://github.com/danielmiessler/fabric#breaking-problems-into-components)
    *   [Too many prompts](https://github.com/danielmiessler/fabric#too-many-prompts)
    *   [The Fabric approach to prompting](https://github.com/danielmiessler/fabric#our-approach-to-prompting)
*   [Installation](https://github.com/danielmiessler/fabric#Installation)
    *   [Migrating](https://github.com/danielmiessler/fabric#Migrating)
    *   [Upgrading](https://github.com/danielmiessler/fabric#Upgrading)
*   [Usage](https://github.com/danielmiessler/fabric#Usage)
*   [Examples](https://github.com/danielmiessler/fabric#examples)
    *   [Just use the Patterns](https://github.com/danielmiessler/fabric#just-use-the-patterns)
*   [Custom Patterns](https://github.com/danielmiessler/fabric#custom-patterns)
*   [Helper Apps](https://github.com/danielmiessler/fabric#helper-apps)
*   [Meta](https://github.com/danielmiessler/fabric#meta)
    *   [Primary contributors](https://github.com/danielmiessler/fabric#primary-contributors)

  

Note

August 20, 2024 — We have migrated to Go, and the transition has been pretty smooth! The biggest thing to know is that **the previous installation instructions in the various Fabric videos out there will no longer work** because they were for the legacy (Python) version. Check the new [install instructions](https://github.com/danielmiessler/fabric#Installation) below.

Intro videos
------------

[](https://github.com/danielmiessler/fabric#intro-videos)

Keep in mind that many of these were recorded when Fabric was Python-based, so remember to use the current [install instructions](https://github.com/danielmiessler/fabric#Installation) below.

*   [Network Chuck](https://www.youtube.com/watch?v=UbDyjIIGaxQ)
*   [David Bombal](https://www.youtube.com/watch?v=vF-MQmVxnCs)
*   [My Own Intro to the Tool](https://www.youtube.com/watch?v=wPEyyigh10g)
*   [More Fabric YouTube Videos](https://www.youtube.com/results?search_query=fabric+ai)

What and why
------------

[](https://github.com/danielmiessler/fabric#what-and-why)

Since the start of 2023 and GenAI we've seen a massive number of AI applications for accomplishing tasks. It's powerful, but _it's not easy to integrate this functionality into our lives._

#### In other words, AI doesn't have a capabilities problem—it has an _integration_ problem.

[](https://github.com/danielmiessler/fabric#in-other-words-ai-doesnt-have-a-capabilities-problemit-has-an-integration-problem)

Fabric was created to address this by enabling everyone to granularly apply AI to everyday challenges.

Philosophy
----------

[](https://github.com/danielmiessler/fabric#philosophy)

> AI isn't a thing; it's a _magnifier_ of a thing. And that thing is **human creativity**.

We believe the purpose of technology is to help humans flourish, so when we talk about AI we start with the **human** problems we want to solve.

### Breaking problems into components

[](https://github.com/danielmiessler/fabric#breaking-problems-into-components)

Our approach is to break problems into individual pieces (see below) and then apply AI to them one at a time. See below for some examples.

[![Image 6: augmented_challenges](https://private-user-images.githubusercontent.com/50654/302028537-31997394-85a9-40c2-879b-b347e4701f06.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MjUwOTAzNjcsIm5iZiI6MTcyNTA5MDA2NywicGF0aCI6Ii81MDY1NC8zMDIwMjg1MzctMzE5OTczOTQtODVhOS00MGMyLTg3OWItYjM0N2U0NzAxZjA2LnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDA4MzElMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQwODMxVDA3NDEwN1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPThmOTJkNTM4OGNlYmNjOTFlZDFjNzM2MzI4OTQ0MGQyZmIzNjdkODQxYTY4OTYzZWEwNWM3OGFhMTNkMDE0ZWYmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.NdBCSamP9pTz9LZo6QyNeOoAfZACwoM5GhFxt0w6eVI)](https://private-user-images.githubusercontent.com/50654/302028537-31997394-85a9-40c2-879b-b347e4701f06.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MjUwOTAzNjcsIm5iZiI6MTcyNTA5MDA2NywicGF0aCI6Ii81MDY1NC8zMDIwMjg1MzctMzE5OTczOTQtODVhOS00MGMyLTg3OWItYjM0N2U0NzAxZjA2LnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDA4MzElMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQwODMxVDA3NDEwN1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPThmOTJkNTM4OGNlYmNjOTFlZDFjNzM2MzI4OTQ0MGQyZmIzNjdkODQxYTY4OTYzZWEwNWM3OGFhMTNkMDE0ZWYmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.NdBCSamP9pTz9LZo6QyNeOoAfZACwoM5GhFxt0w6eVI)

### Too many prompts

[](https://github.com/danielmiessler/fabric#too-many-prompts)

Prompts are good for this, but the biggest challenge I faced in 2023——which still exists today—is **the sheer number of AI prompts out there**. We all have prompts that are useful, but it's hard to discover new ones, know if they are good or not, _and manage different versions of the ones we like_.

One of `fabric`'s primary features is helping people collect and integrate prompts, which we call _Patterns_, into various parts of their lives.

Fabric has Patterns for all sorts of life and work activities, including:

*   Extracting the most interesting parts of YouTube videos and podcasts
*   Writing an essay in your own voice with just an idea as an input
*   Summarizing opaque academic papers
*   Creating perfectly matched AI art prompts for a piece of writing
*   Rating the quality of content to see if you want to read/watch the whole thing
*   Getting summaries of long, boring content
*   Explaining code to you
*   Turning bad documentation into usable documentation
*   Creating social media posts from any content input
*   And a million more…

Installation
------------

[](https://github.com/danielmiessler/fabric#installation)

To install Fabric, [make sure Go is installed](https://go.dev/doc/install), and then run the following command.

# Install Fabric directly from the repo
go install github.com/danielmiessler/fabric@latest

# Run the setup to set up your directories and keys
fabric --setup

### Environment Variables

[](https://github.com/danielmiessler/fabric#environment-variables)

If everything works you are good to go, but you may need to set some environment variables in your `~/.bashrc` or `~/.zshrc` file. Here is an example of what you can add:

# Golang environment variables
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$HOME/.local/bin:$PATH:

### Migration

[](https://github.com/danielmiessler/fabric#migration)

If you have the Legacy (Python) version installed and want to migrate to the Go version, here's how you do it. It's basically two steps: 1) uninstall the Python version, and 2) install the Go version.

# Uninstall Legacy Fabric
pipx uninstall fabric

# Clear any old Fabric aliases
(check your .bashrc, .zshrc, etc.)
# Install the Go version
go install github.com/danielmiessler/fabric@latest
# Run setup for the new version. Important because things have changed
fabric --setup

Then [set your environmental variables](https://github.com/danielmiessler/fabric#environmental-variables) as shown above.

### Upgrading

[](https://github.com/danielmiessler/fabric#upgrading)

The great thing about Go is that it's super easy to upgrade. Just run the same command you used to install it in the first place and you'll always get the latest version.

go install github.com/danielmiessler/fabric@latest

Usage
-----

[](https://github.com/danielmiessler/fabric#usage)

Once you have it all set up, here's how to use it.

usage: fabric -h
Usage:
  fabric \[OPTIONS\]

Application Options:
  -p, --pattern=          Choose a pattern
  -C, --context=          Choose a context
      --session=          Choose a session
  -S, --setup             Run setup
  -t, --temperature=      Set temperature (default: 0.7)
  -T, --topp=             Set top P (default: 0.9)
  -s, --stream            Stream
  -P, --presencepenalty=  Set presence penalty (default: 0.0)
  -F, --frequencypenalty= Set frequency penalty (default: 0.0)
  -l, --listpatterns      List all patterns
  -L, --listmodels        List all available models
  -x, --listcontexts      List all contexts
  -X, --listsessions      List all sessions
  -U, --updatepatterns    Update patterns
  -c, --copy              Copy to clipboard
  -m, --model=            Choose model
  -u, --url=              Choose ollama url (default: http://127.0.0.1:11434)
  -o, --output=           Output to file
  -n, --latest=           Number of latest patterns to list (default: 0)

Help Options:
  -h, --help              Show this help message

Our approach to prompting
-------------------------

[](https://github.com/danielmiessler/fabric#our-approach-to-prompting)

Fabric _Patterns_ are different than most prompts you'll see.

*   **First, we use `Markdown` to help ensure maximum readability and editability**. This not only helps the creator make a good one, but also anyone who wants to deeply understand what it does. _Importantly, this also includes the AI you're sending it to!_

Here's an example of a Fabric Pattern.

https://github.com/danielmiessler/fabric/blob/main/patterns/extract\_wisdom/system.md

[![Image 7: pattern-example](https://private-user-images.githubusercontent.com/50654/302031520-b910c551-9263-405f-9735-71ca69bbab6d.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MjUwOTAzNjcsIm5iZiI6MTcyNTA5MDA2NywicGF0aCI6Ii81MDY1NC8zMDIwMzE1MjAtYjkxMGM1NTEtOTI2My00MDVmLTk3MzUtNzFjYTY5YmJhYjZkLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDA4MzElMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQwODMxVDA3NDEwN1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTYxZTk4NWU5ODIyODdlM2IxNTgzNTVlODk4MmZhODMzZjNiNzQ4NTI1Zjk0ODk2OGY0MzQwNzIzMjk5OTQ0YWEmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.6FjCVYcMNv9G31gQusus915gzvV3KBCphXx9o-l3DzE)](https://private-user-images.githubusercontent.com/50654/302031520-b910c551-9263-405f-9735-71ca69bbab6d.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MjUwOTAzNjcsIm5iZiI6MTcyNTA5MDA2NywicGF0aCI6Ii81MDY1NC8zMDIwMzE1MjAtYjkxMGM1NTEtOTI2My00MDVmLTk3MzUtNzFjYTY5YmJhYjZkLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDA4MzElMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQwODMxVDA3NDEwN1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTYxZTk4NWU5ODIyODdlM2IxNTgzNTVlODk4MmZhODMzZjNiNzQ4NTI1Zjk0ODk2OGY0MzQwNzIzMjk5OTQ0YWEmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.6FjCVYcMNv9G31gQusus915gzvV3KBCphXx9o-l3DzE)

*   **Next, we are extremely clear in our instructions**, and we use the Markdown structure to emphasize what we want the AI to do, and in what order.
    
*   **And finally, we tend to use the System section of the prompt almost exclusively**. In over a year of being heads-down with this stuff, we've just seen more efficacy from doing that. If that changes, or we're shown data that says otherwise, we will adjust.
    

Examples
--------

[](https://github.com/danielmiessler/fabric#examples)

Now let's look at some things you can do with Fabric.

1.  Run the `summarize` Pattern based on input from `stdin`. In this case, the body of an article.

pbpaste | fabric --pattern summarize

2.  Run the `analyze_claims` Pattern with the `--stream` option to get immediate and streaming results.

pbpaste | fabric --stream --pattern analyze\_claims

3.  Run the `extract_wisdom` Pattern with the `--stream` option to get immediate and streaming results from any Youtube video (much like in the original introduction video).

yt --transcript https://youtube.com/watch?v=uXs-zPc63kM | fabric --stream --pattern extract\_wisdom

4.  Create patterns- you must create a .md file with the pattern and save it to ~/.config/fabric/patterns/\[yourpatternname\].

Just use the Patterns
---------------------

[](https://github.com/danielmiessler/fabric#just-use-the-patterns)

[![Image 8: fabric-patterns-screenshot](https://private-user-images.githubusercontent.com/50654/301807224-9186a044-652b-4673-89f7-71cf066f32d8.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MjUwOTAzNjcsIm5iZiI6MTcyNTA5MDA2NywicGF0aCI6Ii81MDY1NC8zMDE4MDcyMjQtOTE4NmEwNDQtNjUyYi00NjczLTg5ZjctNzFjZjA2NmYzMmQ4LnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDA4MzElMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQwODMxVDA3NDEwN1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTE3MTMyNTIwMzcyZjIwYjIxOGY0YTg0ZmUwNTg3OTliNzdkZmZkNDhiNWRkNjQ1NjBmM2IzNTFkZjM5ODI1MTAmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.QW7Cx45h4I8-5cZj_qnhQcL7E4lQItcsoJRCslTvZ-o)](https://private-user-images.githubusercontent.com/50654/301807224-9186a044-652b-4673-89f7-71cf066f32d8.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MjUwOTAzNjcsIm5iZiI6MTcyNTA5MDA2NywicGF0aCI6Ii81MDY1NC8zMDE4MDcyMjQtOTE4NmEwNDQtNjUyYi00NjczLTg5ZjctNzFjZjA2NmYzMmQ4LnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDA4MzElMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQwODMxVDA3NDEwN1omWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTE3MTMyNTIwMzcyZjIwYjIxOGY0YTg0ZmUwNTg3OTliNzdkZmZkNDhiNWRkNjQ1NjBmM2IzNTFkZjM5ODI1MTAmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JmFjdG9yX2lkPTAma2V5X2lkPTAmcmVwb19pZD0wIn0.QW7Cx45h4I8-5cZj_qnhQcL7E4lQItcsoJRCslTvZ-o)If you're not looking to do anything fancy, and you just want a lot of great prompts, you can navigate to the [`/patterns`](https://github.com/danielmiessler/fabric/tree/main/patterns) directory and start exploring!

We hope that if you used nothing else from Fabric, the Patterns by themselves will make the project useful.

You can use any of the Patterns you see there in any AI application that you have, whether that's ChatGPT or some other app or website. Our plan and prediction is that people will soon be sharing many more than those we've published, and they will be way better than ours.

The wisdom of crowds for the win.

Custom Patterns
---------------

[](https://github.com/danielmiessler/fabric#custom-patterns)

You may want to use Fabric to create your own custom Patterns—but not share them with others. No problem!

Just make a directory in `~/.config/custompatterns/` (or wherever) and put your `.md` files in there.

When you're ready to use them, copy them into:

```
~/.config/fabric/patterns/
```

You can then use them like any other Patterns, but they won't be public unless you explicitly submit them as Pull Requests to the Fabric project. So don't worry—they're private to you.

This feature works with all openai and ollama models but does NOT work with claude. You can specify your model with the -m flag

Helper Apps
-----------

[](https://github.com/danielmiessler/fabric#helper-apps)

Fabric also makes use of some core helper apps (tools) to make it easier to integrate with your various workflows. Here are some examples:

`yt` is a helper command that extracts the transcript from a YouTube video. You can use it like this:

yt https://www.youtube.com/watch?v=lQVcbY52\_gY

This will return the transcript from the video, which you can then pipe into Fabric like this:

yt https://www.youtube.com/watch?v=lQVcbY52\_gY | fabric --pattern extract\_wisdom

### `yt` Installation

[](https://github.com/danielmiessler/fabric#yt-installation)

To install `yt`, install it the same way as you install Fabric, just with a different repo name.

go install github.com/danielmiessler/yt@latest

Be sure to add your `YOUTUBE_API_KEY` to `~/.config/fabric/.env`.

Meta
----

[](https://github.com/danielmiessler/fabric#meta)

Note

Special thanks to the following people for their inspiration and contributions!

*   _Jonathan Dunn_ for being the absolute MVP dev on the project, including spearheading the new Go version, as well as the GUI! All this while also being a full-time medical doctor!
*   _Caleb Sima_ for pushing me over the edge of whether to make this a public project or not.
*   _Eugen Eisler_ and _Frederick Ros_ for their invaluable contributions to the Go version
*   _Joel Parish_ for super useful input on the project's Github directory structure..
*   _Joseph Thacker_ for the idea of a `-c` context flag that adds pre-created context in the `./config/fabric/` directory to all Pattern queries.
*   _Jason Haddix_ for the idea of a stitch (chained Pattern) to filter content using a local model before sending on to a cloud model, i.e., cleaning customer data using `llama2` before sending on to `gpt-4` for analysis.
*   _Andre Guerra_ for assisting with numerous components to make things simpler and more maintainable.

### Primary contributors

[](https://github.com/danielmiessler/fabric#primary-contributors)

[![Image 9](https://avatars.githubusercontent.com/u/50654?v=4)](https://github.com/danielmiessler) [![Image 10](https://avatars.githubusercontent.com/u/9218431?v=4)](https://github.com/xssdoctor) [![Image 11](https://avatars.githubusercontent.com/u/688589?v=4)](https://github.com/sbehrens) [![Image 12](https://avatars.githubusercontent.com/u/10410523?v=4)](https://github.com/agu3rra)


[2] Title: Why I Created Fabric
[2] URL Source: https://danielmiessler.com/p/fabric-origin-story
[2] Description: An overview of the AI workflows I built in 2023, and what became the <strong>Fabric</strong> project.
[2] Markdown Content:
Why I Created Fabric
===============
 

[![Image 1: Unsupervised Learning logo](https://media.beehiiv.com/cdn-cgi/image/fit=scale-down,format=auto,onerror=redirect,quality=80/uploads/publication/logo/6af6184f-358e-47dc-8501-ddfd3b0a8928/thumb_small-blue-trans-back.png) Unsupervised Learning](https://danielmiessler.com/)

Categories

[Newsletter](https://danielmiessler.com/t/Newsletter)[Podcast](https://omny.fm/shows/unsupervised-learning)[About](https://danielmiessler.com/c/about)[Become a Member](https://danielmiessler.com/upgrade/)[Member Portal](https://danielmiessler.com/p/ul-member-portal)[Support](https://danielmiessler.com/c/support)

Login[Subscribe](https://danielmiessler.com/subscribe)

0

*   [Unsupervised Learning](https://danielmiessler.com/)
*   Posts
*   Why I Created Fabric

      

Why I Created Fabric
====================

An overview of the AI workflows I built in 2023, and what became the Fabric project.
------------------------------------------------------------------------------------

![Image 2: Author](https://media.beehiiv.com/cdn-cgi/image/fit=scale-down,format=auto,onerror=redirect,quality=80/uploads/user/profile_picture/75870627-4c61-4be1-8969-2bbb164fda67/thumb_daniel-black-background.png)

[Daniel Miessler](https://danielmiessler.com/authors/75870627-4c61-4be1-8969-2bbb164fda67)  
February 01, 2024

[](https://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Fdanielmiessler.com%2Fp%2Ffabric-origin-story)[](https://twitter.com/intent/tweet?text=An+overview+of+the+AI+workflows+I+built+in+2023%2C+and+what+became+the+Fabric+project.&url=https%3A%2F%2Fdanielmiessler.com%2Fp%2Ffabric-origin-story&via=danielmiessler)[](https://www.linkedin.com/sharing/share-offsite?url=https%3A%2F%2Fdanielmiessler.com%2Fp%2Ffabric-origin-story)

In this video on [David Bombal’s podcast](https://www.youtube.com/@davidbombal?utm_source=danielmiessler.com&utm_medium=referral&utm_campaign=why-i-created-fabric), I talk through the AI tooling I spent 2023 building.

[![Image 3](https://media.beehiiv.com/cdn-cgi/image/fit=scale-down,format=auto,onerror=redirect,quality=80/uploads/asset/file/bff9f639-897b-4122-96c7-68f276c34798/image.png?t=1706845986)](https://youtu.be/vF-MQmVxnCs?utm_source=danielmiessler.com&utm_medium=referral&utm_campaign=why-i-created-fabric)

Click to watch the video on YouTube

The video covers:

*   How I captured my desired outcomes
    
*   How I break everything into components
    
*   How I apply AI to those individual components
    
*   How I call those AI commands from the CLI
    
*   **How I chain those commands together to accomplish full workflows!**
    

This tooling is what became [Fabric](https://github.com/danielmiessler/fabric?utm_source=danielmiessler.com&utm_medium=referral&utm_campaign=why-i-created-fabric) in the beginning of 2024.

[![Image 4](https://opengraph.githubassets.com/085617f43d5690cb6f5f93bdcaf4b3ef695a283d2be33e0a96404e96137dda87/danielmiessler/fabric) fabric is an open-source framework for augmenting humans using AI. The goal of the project is to provide a universally accessible layer of AI that anyone can use to enhance their life or work. github.com/danielmiessler/fabric](https://github.com/danielmiessler/fabric?utm_source=danielmiessler.com&utm_medium=referral&utm_campaign=why-i-created-fabric)

![Image 5: SECURITY | AI | MEANING :: One security-minded AI builder's continuous stream of original ideas, analysis, tools, and mental models on how to build a successful and meaningful life in a world full of AI.](https://media.beehiiv.com/cdn-cgi/image/fit=scale-down,format=auto,onerror=redirect,quality=80/uploads/publication/logo/6af6184f-358e-47dc-8501-ddfd3b0a8928/thumb_small-blue-trans-back.png)Unsupervised Learning

SECURITY | AI | MEANING :: One security-minded AI builder's continuous stream of original ideas, analysis, tools, and mental models on how to build a successful and meaningful life in a world full of AI.

Home

[Posts](https://danielmiessler.com/)[Authors](https://danielmiessler.com/authors)

Account

[Upgrade](https://danielmiessler.com/upgrade)

Newsletter

[Newsletter](https://danielmiessler.com/t/Newsletter)

[](https://twitter.com/danielmiessler)[](https://www.linkedin.com/in/danielmiessler)[](https://www.youtube.com/channel/UCnCikd0s4i9KoDtaHPlK-JA)[](https://tiktok.com/@danielmiessler)[](https://instagram.com/danielmiessler)[](https://rss.beehiiv.com/feeds/gQxaV1KHkQ.xml)

© 2024 Unsupervised Learning.

[Privacy Policy](https://beehiiv.com/privacy)[Terms of Use](https://beehiiv.com/tou)

[Powered by beehiiv](https://www.beehiiv.com/?utm_source=UnsupervisedLearning&utm_medium=footer)


[3] Title: Analyzing Threat Reports with Fabric
[3] URL Source: https://danielmiessler.com/p/fabric-pattern-analyze-threat-report
[3] Description: - Build a cybersecurity culture ... the specific Pattern. analyze_threat_report. GitHub - danielmiessler/fabric: fabric is <strong>an open-source framework for augmenting humans using AI</strong>....
[3] Markdown Content:
We’ve just added a new Pattern to `fabric`.

It’s called `analyze_threat_report`, and it’s designed to extract all the most valuable parts of a cybersecurity threat report like the DBIR report, Crowdstrike, Blackberry, etc.

The output (from the Crowdstrike 2024 Global Threat Report)
-----------------------------------------------------------

```
ONE-SENTENCE-SUMMARY:

The 2024 CrowdStrike Global Threat Report highlights the accelerated pace and sophistication of cyberattacks, emphasizing the critical need for advanced, AI-driven cybersecurity measures in the face of evolving threats.

TRENDS:

- Generative AI lowers the entry barrier for cyberattacks, enabling more sophisticated threats.

- Identity-based attacks and social engineering are increasingly central to adversaries' strategies.

- Cloud environments are under greater threat as adversaries advance their capabilities.

- The use of legitimate tools by attackers complicates the detection of malicious activities.

- A significant rise in supply chain attacks, exploiting trusted software for maximum impact.

- The potential targeting of global elections by adversaries to influence geopolitics.

- The emergence of 34 new adversaries, including a newly tracked Egypt-based adversary, WATCHFUL SPHINX.

- A 60% increase in interactive intrusion campaigns observed, with technology sectors being the primary target.

- A notable rise in ransomware and data-theft extortion activities, with a 76% increase in victims named on dedicated leak sites.

- North Korean adversaries focus on financial gain through cryptocurrency theft and intelligence collection.

- Stealth tactics are increasingly employed to evade detection and move laterally within networks.

- Access brokers play a crucial role in providing initial access to eCrime threat actors.

- A shift towards ransomware-free data leak operations among big game hunting adversaries.

- The growing use of cloud-conscious techniques by adversaries to exploit cloud vulnerabilities.

- An increase in the use of legitimate remote monitoring and management tools by eCrime actors.

- The persistence of access brokers in facilitating cyberattacks through advertised accesses.

- Law enforcement's increased focus on disrupting big game hunting operations and their supporting infrastructure.

- The rise of macOS malware variants targeting information stealers to expand eCrime profit opportunities.

- The adaptation of malware delivery techniques following patches for Mark-of-the-Web bypass vulnerabilities.

STATISTICS:

- Cloud-conscious cases increased by 110% year over year (YoY).

- A 76% YoY increase in victims named on eCrime dedicated leak sites.

- 34 new adversaries tracked by CrowdStrike, raising the total to 232.

- Cloud environment intrusions increased by 75% YoY.

- 84% of adversary-attributed cloud-conscious intrusions were focused on eCrime.

- A 60% year-over-year increase in the number of interactive intrusion campaigns observed.

- The average breakout time for interactive eCrime intrusion activity decreased from 84 minutes in 2022 to 62 minutes in 2023.

- The number of accesses advertised by access brokers increased by almost 20% compared to 2022.

- A 583% increase in Kerberoasting attacks in 2023.

QUOTES:

- "You don’t have a malware problem, you have an adversary problem."

- "The speed and ferocity of cyberattacks continue to accelerate."

- "Generative AI has the potential to lower the barrier of entry for low-skilled adversaries."

- "Identity-based attacks take center stage."

- "We are entering an era of a cyber arms race where AI will amplify the impact."

- "The continued exploitation of stolen identity credentials."

- "The growing menace of supply chain attacks."

- "Adversaries are advancing their capabilities to exploit the cloud."

- "The use of legitimate tools to execute an attack impedes the ability to differentiate between normal activity and a breach."

- "Organizations must prioritize protecting identities in 2024."

REFERENCES:

- CrowdStrike Falcon® XDR platform

- CrowdStrike Counter Adversary Operations (CAO)

- CrowdStrike Falcon® Intelligence

- CrowdStrike® Falcon OverWatchTM

- Microsoft Outlook (CVE-2023-23397)

- Azure Key Vault

- CrowdStrike Falcon® Identity Threat Protection

- CrowdStrike Falcon® Fusion Playbooks

- CrowdStrike Falcon® Adversary OverwatchTM

- CrowdStrike Falcon® Adversary Intelligence

- CrowdStrike Falcon® Adversary Hunter

RECOMMENDATIONS:

- Implement phishing-resistant multifactor authentication and extend it to legacy systems and protocols.

- Educate teams on social engineering and implement technology that can detect and correlate threats across identity, endpoint, and cloud environments.

- Implement cloud-native application protection platforms (CNAPPs) for full cloud visibility, including into applications and APIs.

- Gain visibility across the most critical areas of enterprise risk, including identity, cloud, endpoint, and data protection telemetry.

- Drive efficiency by using tools that unify threat detection, investigation, and response in one platform for unrivaled efficiency and speed.

- Build a cybersecurity culture with user awareness programs to combat phishing and related social engineering techniques.
```

The project
-----------

To use this, and all the other Patterns in Fabric, head over to the project page.


[4] Title: How to run Daniel Miesslers Fabric in a KM shell? github link in the body
[4] URL Source: https://forum.keyboardmaestro.com/t/how-to-run-daniel-miesslers-fabric-in-a-km-shell-github-link-in-the-body/35913
[4] Description: Hello everyone, I got the little things working, but not the starting of python programs via the shell, maybe someone knows better and has got <strong>fabric</strong> running. I would like to monitor folders with Keyboard Maestro and if there are files in the folder, execute different patterns automatically ...
[4] Published Time: 2024-04-19T09:48:24+00:00
[4] Markdown Content:
![Image 1: Keyboard Maestro Discourse](blob:https://forum.keyboardmaestro.com/edcd92a17d838b69d51966c5b06a2fc1)

Loading


[5] Title: Releases · danielmiessler/fabric
[5] URL Source: https://github.com/danielmiessler/fabric/releases
[5] Description: fabric is <strong>an open-source framework for augmenting humans using AI</strong>. It provides a modular framework for solving specific problems using a crowdsourced set of AI prompts that can be used anywhere. - Releases · danielmiessler/fabric
[5] Markdown Content:
[Skip to content](https://github.com/danielmiessler/fabric/releases#start-of-content)

