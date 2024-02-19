# An Exercise in Automated Programming (Winter 2024)
## Introduction
This assignment asked students to explore using automated code generation methods to complete an assignment from earlier in the course. The question posed asked to think abotu how AI bots may either completely replace or serve as companions to make coding more efficient and cheaper for companies.
## Automated code generation
I personally found using go generate a little difficult to grasp at first. However, a Medium article from Edouard Tavinor written in 2017, helped me by showing how the command could be used to generate boiler plate code when you have similar structures. It's very cool how you are able to make it function like a combination of generics and superclasses in Java, where the functions and attributes are inherited, but to the specified type in the generate command.
## AI-assisted programming
Comparison data provided to me comes from Anscombe's Quartet as well as Dr. Miller's OLS analysis (written in both Python and R). I utlized Dr Miller's OLS outputs to compare my linear regression results. Unit tests for this application compares function outputs to that from Python's OLS output, including slope, y-intercept, r-squared, adjusted r-squared, and F-statistic.
## AI-generated code
Given my current expertise with Go, I personally would not recommend using Go for statistics. With the recommended packages, I found them insufficient for obtaining necessary statistical values. Even to obtain the slope, I had to code a custom function to perform the calculation, which though possible, took a considerable amount of time. Python and R allows me to obtain this information with minimal coding knowledge. Additionally, much more testing needs to be conducted on these custom functions as I found they worked for some inputs but failed others (see observation y4). At this time, Python and R afford much more statistical functionality and convenience.
### Training Materials
- go:generate documentation: https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source
- jennifer documentation: https://pkg.go.dev/github.com/dave/jennifer#section-documentation
- GitHub Copilot documentation: https://copilot.github.com/docs/
- LogRocket, Using go generate to reduce boilerplate code: https://blog.logrocket.com/using-go-generate-reduce-boilerplate-code/
- Medium, go generate by Edouard Tavinor: https://ehrt74.medium.com/go-generate-89b20a27f7f9
