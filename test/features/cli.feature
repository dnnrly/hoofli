Feature: Simple CLI commands

    @Acceptance
    Scenario: Runs command correctly
        When the app runs with parameters ""
        Then the app exits with an error
        And the app output contains "required parameter --input missing"

    @Acceptance
    Scenario: Prints help correctly
        When the app runs with parameters "-h"
        Then the app exits without error
        And the app output contains "Options:"

    @Acceptance
    Scenario: Reports error for HAR parsing failure
        When the app runs with parameters "--input reference/har/invalid.har"
        Then the app exits with an error
        And the app output contains "unable to parse HAR file"

    @Acceptance
    Scenario: Generates plantiml to STDOUT from a HAR file
        When the app runs with parameters "--input reference/har/google-frontpage.har"
        Then the app exits without error
        And the app output contains "->Browser : Google"
        And the app output contains "Browser->\"www.google.com\" ++ : GET https://www.google.com/"
        And the app output contains "gstatic"

    @Acceptance
    Scenario: Excludes interactions by URL
        When the app runs with parameters "--input reference/har/google-frontpage.har --exclude-url (adsense|gstatic|doubleclick|apis.google|ogs.google)"
        Then the app exits without error
        And the app output does not contain "gstatic"

    @Acceptance
    Scenario: Excludes interactions by header value
        When the app runs with parameters "--input reference/har/google-frontpage.har --exclude-header content-type=image/.+"
        Then the app exits without error
        And the app output does not contain "googlelogo_color_272x92dp.png"
