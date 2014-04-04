#
Feature: Update payment
	In order for users to add sub items to the payment it must be updated.
	Sub items help communicate to the client what they are paying for.
	They also help keep track of what work has been completed.

	Scenario: Update
		Given an existing set of payments
		When I update the payment items
		Then a payment is returned with the same sub items
