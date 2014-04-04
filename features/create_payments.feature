#
Feature: Create payments
	In order for users to create an agreement there need to be terms for payment.
	Creating payments allows freelancers to communicate with their about what the expectations
	are around when they will get paid, how much and for what services.

	Scenario: Create and Save
		Given a new set of payments
		When I save them
		Then they each have an id
