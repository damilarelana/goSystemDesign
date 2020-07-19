IMAGE_NAME = "ubuntu/focal64" # os image for vm i.e. latest ubuntu focal [which has virtualbox pre-installed]
N = 2 # num of nodes in cluster i.e. excluding the master3 nodes [1 master + 2 workers]

Vagrant.configure("2") do |config| # defines Vagrant configuration version 2
    config.ssh.insert_key = false

    config.vm.provider "virtualbox" do |v|
        v.memory = 2048
        v.cpus = 2
    end
      
    config.vm.define "k8s-master" do |master|
        master.vm.box = IMAGE_NAME
        master.vm.network "private_network", ip: "192.168.15.10" # '192.168.15.10' assumes it's same network of virtual machine hosts, and also available
        master.vm.hostname = "k8s-master"
        master.vm.provision "ansible" do |ansible|
            ansible.playbook = "kubernetes-setup/master-playbook.yml"
            ansible.extra_vars = {
                node_ip: "192.168.15.10",
            }
        end
    end

    (1..N).each do |i| # iterates to create clones of the same number of Virtualboxs
        config.vm.define "k8s-node-#{i}" do |node|
            node.vm.box = IMAGE_NAME
            node.vm.network "private_network", ip: "192.168.15.#{i + 10}"
            node.vm.hostname = "k8s-node-#{i}"
            node.vm.provision "ansible" do |ansible|
                ansible.playbook = "kubernetes-setup/node-playbook.yml"
                ansible.extra_vars = {
                    node_ip: "192.168.15.#{i + 10}",
                }
            end
        end
    end
