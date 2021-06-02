Feature: service hello
  In order to be happy
  As a developer
  I need to be able to create a service

  Scenario: Print "Hello, World!"
    Given a service
    When I start it
    Then it should display "Hello, World!" message
