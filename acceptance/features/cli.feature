Feature: Simple CLI commands

    @Acceptance
    Scenario: Runs command correctly
        When the app runs with parameters ""
        Then the app exits without error