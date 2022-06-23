variable "primary_hosted_zone" {
  type        = string
  description = "Primary hosted zone for this Virtual Gateway. Populate `TF_VAR_primary_hosted_zone` before running any tests to have this value populated."
}
