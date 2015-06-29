# -*- mode: ruby -*-
# vi: set ft=ruby :

hosts = {
  "proc_dump" => {"box" => "ubuntu-14.04"}
}


Vagrant.configure("2") do |config|
  hosts.each do |name, settings|
    config.vm.synced_folder "./", "/vagrant"
    
    config.vm.define name do |machine|
      machine.vm.box = settings.fetch("box")
    end
  end
end
