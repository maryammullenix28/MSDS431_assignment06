# An Exercise in Automated Programming (Winter 2024)
## Introduction
This assignment asked students to explore using automated code generation methods to complete an assignment from earlier in the course. The question posed asked to think abotu how AI bots may either completely replace or serve as companions to make coding more efficient and cheaper for companies.
## Automated code generation
I personally found using go generate a little difficult to grasp at first. However, a Medium article from Edouard Tavinor written in 2017, helped me by showing how the command could be used to generate boiler plate code when you have similar structures. It's very cool how you are able to make it function like a combination of generics and superclasses in Java, where the functions and attributes are inherited, but to the specified type in the generate command.
## AI-assisted programming
As for using GitHub Copilot, I found the tool incredibly useful at creating more efficient code and coming up with checks and solutions for edge cases. For example, in my function createCoordinateSlice in stats_go.go, I assumed that the x and y variables would be the same length. However, when I asked GitHub Copilot to optimize my code, it made the suggestion to check that two were the same length, and print an error if not. This shows that the tool could help with developing more comprehensive tests and think of edge cases, especially for junior programmers who have less experience in this. Some may argue that it takes the thinking out of coding, but it's difficult for junior programmers to be aware of what they don't know. Now, with that example, it's likely that I will consider this scenario in the future for similar functions involving correlated slices.
## AI-generated code
I chose to utilize ChatGPT to generate code for the Testing Go for Statistics assignment. To do so, I provided ChatGPT with two key pieces of information, the inputs (x and y variables) and the intended output (a statiscal print out of key information, like R-square, F-stat, etc.). Surprisingly, ChatGPT actually provided fully functional code without any errors. Though I've used ChatGPT to assist me with coding in the past, it's usually very general or to provide examples on how a certain function or data structure works. Even these simple requests lead to many errors and having to personally tweak the code to make it run. That being said, the program generated provided incorrect outputs that did not align with my original assignment's output nor the R or Python outputs I compared mine to. When made aware that the calculations could be incorrect, ChatGPT insisted that it was just rounding error. 
## Recommendation

### Training Materials
- go:generate documentation: https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source
- jennifer documentation: https://pkg.go.dev/github.com/dave/jennifer#section-documentation
- GitHub Copilot documentation: https://copilot.github.com/docs/
- LogRocket, Using go generate to reduce boilerplate code: https://blog.logrocket.com/using-go-generate-reduce-boilerplate-code/
- Medium, go generate by Edouard Tavinor: https://ehrt74.medium.com/go-generate-89b20a27f7f9
