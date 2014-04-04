#
Feature: Update last action
	In order for to know where an agreement stands in terms of progress
	we need to keep track of the most current actions taken. In order to keep track we must be able to update the last action of a payment.

	Scenario: Update
		Given an existing set of payments
		When I take one of them
		When I create a new action for that payment
		Then an action is returned