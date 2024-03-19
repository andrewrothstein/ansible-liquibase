output "aws_instances" {
  value = [for instance in aws_instance.this : instance.public_ip]
}
