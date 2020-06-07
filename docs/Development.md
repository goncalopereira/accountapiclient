# Development

## Spike
First attempt was to add the client with all the requests and display responses into the available integration API

## Marshal
From there I was able to add the JSON marshal/unmarshal into the correct types and test that

## Delete
Had some bash scripts to generate data based on the Form3 website so was able to check for correct behavior
Implemented Delete, unit tests and Request to wrap http.client for testing

## Fetch
Built Fetch and refactored the API config/URL generation into all endpoints

## Create
Built Create and refactored

## List 
Built List and refactored

## Refactor Output
With all commands I could refactor JSON marshal/unmarshal to be cleaner as well as the Output type
Divided between having three return types (Account,ErrorMessage,error) and two (IOutput, error)

## Refactor Accounts
Refactored the Accounts and standard err to be more standardised. Always a lot of shuffling around as I'm new to Go!

## Refactor Request
Could refactor and clean up Request to a single method
Added the NoOp type to replace a nil option (Null Object)

## Go idiomatic changes
Some of the changes that seem too OOP for Go was the wrapper for JSON library for defensive coding against IO that I removed.
As well as standardising error names to err.
Could have optionally also removed the http wrapper and use httptest on every command unit test.

## Testing extras
Finishing up e2e tests which run after unit tests. 
Trying to grab edge cases against each endpoint but not all features actually work on Fake API.
According to Go documentation also moved common test libraries to internal/ instead of test/


  
