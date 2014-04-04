#
Feature: Fetch payments
	In order for users to see a full agreement or to interact make any payment related actions
	the user must be able to fetch and see the payments associated with an agreement.

	Scenario: Fetch By Version ID
		Given an existing set of payments
		When I fetch them by the agreements version ID
		Then at least one is returned
