Feature: Validate Process Payment
  Validate the Payment for processing

  #Scenario: Validate a Payment with zero amount
  #Scenario: Validate a Payment with negative amount
  #Scenario: Validate a Payment with amount more than 5000
  #Scenario: Validate a Payment with empty Type
  Scenario: Validate a Payment with Type CH1
    Given a New Payment
    And the payment has amount 50
    And the payment has status "pending"
    And the payment has type "CH1"
    When validate process the payment
    Then returns a non error
    And returns a payment with detail "to_review"

  #Scenario: Validate a Valid Pending Payment
  #  Given a New Payment
  #  And the payment has amount 50
  #  And the payment has status "pending"
  #  When process the payment
  #  Then returns a non error
  #  And returns a payment with status "pending"
