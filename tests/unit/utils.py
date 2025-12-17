import os
import re
import subprocess

class TerraformWrapper:
    def __init__(self, terraform_bin_path=None):
        self.terraform_bin_path = terraform_bin_path or 'terraform'

    def init(self):
        try:
            subprocess.check_output([self.terraform_bin_path, 'init'])
        except subprocess.CalledProcessError as e:
            raise Exception(f"Failed to run terraform init: {e}")

    def apply(self):
        try:
            subprocess.check_output([self.terraform_bin_path, 'apply', '-auto-approve'])
        except subprocess.CalledProcessError as e:
            raise Exception(f"Failed to run terraform apply: {e}")

    def destroy(self):
        try:
            subprocess.check_output([self.terraform_bin_path, 'destroy', '-auto-approve'])
        except subprocess.CalledProcessError as e:
            raise Exception(f"Failed to run terraform destroy: {e}")

    def validate(self):
        try:
            subprocess.check_output([self.terraform_bin_path, 'validate'])
        except subprocess.CalledProcessError as e:
            raise Exception(f"Failed to run terraform validate: {e}")

    def show(self, instance_name):
        try:
            output = subprocess.check_output([self.terraform_bin_path, 'show', instance_name])
            return output.decode('utf-8')
        except subprocess.CalledProcessError as e:
            raise Exception(f"Failed to run terraform show: {e}")