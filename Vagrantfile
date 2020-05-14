# -*- mode: ruby -*-
# vi: set ft=ruby :

require 'yaml'

begin
  current_dir    = File.dirname(File.expand_path(__FILE__))
  # config.yml is ignored by git, i.e., .gitignore
  configs        = YAML.load_file("#{current_dir}/config.yml")
  vagrant_config = configs['configs'][configs['configs']['use']]
rescue StandardError => msg
  vagrant_config = {}
end

$welcome_msg = <<MSG
-----------------------------------------------------------------
S-ENDA Prototype started from the latest development version from
https://hub.docker.com/repository/docker/metno/senda-prototype

URLs and ports:
 - catalog-service-api  - http://10.20.30.10:80
 - dynamic-geoasset-api - http://10.20.30.10:8080
 - data-dashboard       - http://10.20.30.10:8081
 - websocket-bridge     - 10.20.30.10:8084

To build/re-build and start services from source run and re-run:

  BUILD=1 vagrant up

Stop and remove development environment with:

  vagrant destroy -f
-----------------------------------------------------------------
MSG

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/bionic64"
  config.vm.box_check_update = false

  config.vm.network "private_network", ip: "10.20.30.10"

  config.vm.provider "virtualbox" do |vb|
    vb.memory = "4096"
    vb.cpus = 4
    vb.default_nic_type = "virtio"
  end

  config.vm.define "default" do |config|
    if vagrant_config != {}
      config.vm.network "public_network", ip: vagrant_config['ip'], netmask: vagrant_config['netmask'], bridge: vagrant_config['bridge']
      config.vm.provision "shell", run: "always", inline: "ip route add default via #{ vagrant_config['gateway'] } metric 10 || exit 0"
      config.vm.hostname = vagrant_config['hostname']
    end
  end

  config.vm.provision "00-bootstrap", type: "shell", inline: <<-SHELL
    apt-get update
    apt-get install -y git webhook wget unattended-upgrades
    apt-get install -y docker.io docker-compose
    cp /vagrant/hooks.json /etc/webhook.conf
    systemctl restart webhooks
  SHELL

  config.vm.provision "50-rebuild", type: "shell", run: "always", inline: <<-SHELL
    cd /vagrant
    if [[ -n "#{ENV['BUILD']}" ]]
    then
      docker-compose -f docker-compose.yml -f docker-compose.build.yml build --pull
    fi
    docker-compose up -d
  SHELL

  config.vm.post_up_message = $welcome_msg
end
