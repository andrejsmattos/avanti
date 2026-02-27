# Gera uma chave privada RSA de 4096 bits
resource "tls_private_key" "rsa_key" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

# Cria o Key Pair na AWS usando a chave pública gerada
resource "aws_key_pair" "ec2_key_pair" {
  key_name   = "ec2-instance-key"
  public_key = tls_private_key.rsa_key.public_key_openssh
}

# Salva a chave privada no filesystem Linux do WSL (~/.ssh)
resource "local_sensitive_file" "private_key_pem" {
  filename             = pathexpand("~/.ssh/ec2-instance-key.pem")
  content              = tls_private_key.rsa_key.private_key_pem
  file_permission      = "0400"
  directory_permission = "0700"
}