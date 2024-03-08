package models

type Analytics struct {
	Enabled bool
	//Whether or not analytics is enabled.

	Sampling_type string
	//The type of the sampling

	/*Allowed values:
	*PROBABILITY
	*TEMPORAL
	*COUNT
	 */

	Sampling_value string
	//The value of the sampling

	Logging_condition               string
	Logging_messageCondition        string
	Logging_content_headers         bool
	Logging_content_messageHeaders  bool
	Logging_content_payload         bool
	Logging_content_messagePayload  bool
	Logging_content_messageMetadata bool
	Logging_phaserequest            bool
	Logging_phaseresponse           bool
	Logging_mode_endpoint           bool
	Logging_mode_entrypoint         bool
}
