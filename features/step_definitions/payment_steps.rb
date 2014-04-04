#encoding: utf-8
require "rubygems"
require "majordomo"
require "json"
require "securerandom"

Before do
	@client = Majordomo::Client.new("tcp://localhost:5555", false)
end

After do
	@client.close
end

Given /^a new set of payments$/ do
	@versionID = SecureRandom.uuid
	@payments = [{:title=>"1", :versionID => @versionID}, {:title=>"2", :versionID => @versionID}]
end

Given /^an existing set of payments$/ do
	@versionID = SecureRandom.uuid
	payments = [{:title=>"1", :versionID => @versionID}, {:title=>"2", :versionID => @versionID}]
	jsonString = [JSON.generate(payments)].pack('m0')
	reply = @client.send_and_receive("Payments", '{"Method":"POST", "Path":"/agreements/v/'+@versionID+'/payments", "Body":"'+jsonString+'"}')
	@savedPayments = JSON.parse(JSON.parse(reply[0])["body"].unpack('m0')[0])
end

When /^I save them$/ do
	jsonString = [JSON.generate(@payments)].pack('m0')
	reply = @client.send_and_receive("Payments", '{"Method":"POST", "Path":"/agreements/v/'+@versionID+'/payments", "Body":"'+jsonString+'"}')
	@savedPayments = JSON.parse(JSON.parse(reply[0])["body"].unpack('m0')[0])
end

When /^I fetch them by the agreements version ID$/ do
	reply = @client.send_and_receive("Payments", '{"Method":"GET", "Path":"/agreements/v/'+ @versionID+'/payments"}')
	@savedPayments = JSON.parse(reply[0])["body"].unpack('m0')[0]
	@savedPayments = JSON.parse(@savedPayments)
end

When /^I update the payment items$/ do
	@updatedPayment = @savedPayments[0]
	@updatedPayment["paymentItems"] = [{:title => "test1"}]
	jsonString = [JSON.generate(@updatedPayment)].pack('m0')
	reply = @client.send_and_receive("Payments", '{"Method":"PUT", "Path":"/payments/'+@updatedPayment["id"]+'", "Body":"'+jsonString+'"}')
	@savedPayment = JSON.parse(JSON.parse(reply[0])["body"].unpack('m0')[0])
end

When /^I take one of them$/ do
	@savedPayment = @savedPayments[0]
end

When /^I create a new action for that payment$/ do
	jsonString = [JSON.generate({:name => "completed"})].pack('m0')
	reply = @client.send_and_receive("Payments", '{"Method":"POST", "Path":"/payments/'+@savedPayment["id"]+'/action", "Body":"'+jsonString+'"}')
	@action = JSON.parse(JSON.parse(reply[0])["body"].unpack('m0')[0])
end

Then /^they each have an id$/ do
	@savedPayments.should be_a_kind_of(Array)

	@savedPayments.each {|x| 
		(!x["id"].nil? or (!x["id"].nil? and !x["id"].empty?)).should be_true
	}
end

Then /^at least one is returned$/ do
	@savedPayments.should be_a_kind_of(Array)
end

Then /^a payment is returned with the same sub items$/ do
	@savedPayment["paymentItems"].should be_a_kind_of(Array)
	@savedPayment["paymentItems"].length.should eql(@updatedPayment["paymentItems"].length)
end

Then /^an action is returned$/ do
	@action.should be_a_kind_of(Hash)
	@action["name"].should_not be_nil
	@action["name"].should_not be_empty
end