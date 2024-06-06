    create logrotate in linux
    /home/$USER/.gas/logs/*.log {
            su root root
            daily
            missingok
            rotate 60
            compress
            create
            copytruncate
            dateext
            notifempty
            dateformat -%Y-%m-%d
            dateyesterday
    }

    sudo logrotate -f -v /etc/logrotate.d/gas
