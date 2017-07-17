#!/bin/bash


jar="/home/bjcheny/selenium/selenium-server-standalone-2.44.0.jar"
if [[ ! -f "${jar}" ]]; then

    # Following the guide found at this page
    # http://programmingarehard.com/2014/03/17/behat-and-selenium-in-vagrant.html

    echo -e "\r\nUpdating system ...\r\n"
    sudo apt-get update


    # Create folder to place selenium in
    #
    echo -e "\r\nCreating folder to place selenium in ...\r\n"
    sudo mkdir ~/selenium
    cd ~/selenium


    # Get Selenium and install headless Java runtime
    #
    echo -e "\r\nInstalling Selenium and headless Java runtime ...\r\n"
    sudo wget http://selenium-release.storage.googleapis.com/2.44/selenium-server-standalone-2.44.0.jar
    sudo apt-get install openjdk-7-jre-headless -y


    # Install Firefox
    #
    echo -e "\r\nInstalling Firefox ...\r\n"
    sudo apt-get install firefox -y


    # Install headless GUI for firefox.  'Xvfb is a display server that performs graphical operations in memory'
    #
    echo -e "\r\nInstalling XVFB (headless GUI for Firefox) ...\r\n"
    sudo apt-get install xvfb -y
fi

# Finally, starting up Selenium server
#
echo -e "\r\nStarting up Selenium server ...\r\n"
DISPLAY=:1 xvfb-run java -jar ~/selenium/selenium-server-standalone-2.44.0.jar
