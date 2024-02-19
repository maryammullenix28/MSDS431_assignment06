# An Exercise in Automated Programming (Winter 2024)
## Introduction
This assignment asked students to explore using automated code generation methods to complete an assignment from earlier in the course. The question posed asked to think abotu how AI bots may either completely replace or serve as companions to make coding more efficient and cheaper for companies.
## Automated code generation
This statistical application performs a linear regression analysis on sets of x and y data points. This was implemented using Go's fmt and time package, montanaflynn's stats package, and gonum's distuv package. Using package functions, I created custom functions that calculates the slope, y-intercept, y-hat, SSE, SSTO, r-squared, adjusted r-squared, and F-statistic.
## AI-assisted programming
Comparison data provided to me comes from Anscombe's Quartet as well as Dr. Miller's OLS analysis (written in both Python and R). I utlized Dr Miller's OLS outputs to compare my linear regression results. Unit tests for this application compares function outputs to that from Python's OLS output, including slope, y-intercept, r-squared, adjusted r-squared, and F-statistic.
## AI-generated code
Given my current expertise with Go, I personally would not recommend using Go for statistics. With the recommended packages, I found them insufficient for obtaining necessary statistical values. Even to obtain the slope, I had to code a custom function to perform the calculation, which though possible, took a considerable amount of time. Python and R allows me to obtain this information with minimal coding knowledge. Additionally, much more testing needs to be conducted on these custom functions as I found they worked for some inputs but failed others (see observation y4). At this time, Python and R afford much more statistical functionality and convenience.
### Training Materials
- go:generate documentation: https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source
- jennifer documentation: https://pkg.go.dev/github.com/dave/jennifer#section-documentation
- GitHub Copilot documentation: https://copilot.github.com/docs/
