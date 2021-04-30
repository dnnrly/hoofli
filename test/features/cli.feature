Feature: Simple CLI commands

    @Acceptance
    Scenario: Runs command correctly
        When the app runs with parameters ""
        Then the app exits with an error
        And the app output contains "must specify HAR input"

    @Acceptance
    Scenario: Generates plantiml to STDOUT from a HAR file
        When the app runs with parameters "-har reference/har/google-frontpage.har"
        Then the app exits without error
        And the app output contains "->Browser : Google"
        # And the app output contains "Browser->www.google.com : GET https://www.google.com/"
