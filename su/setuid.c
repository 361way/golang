#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
#include <pwd.h>

int main(void)
{
    int current_uid = getuid();
    printf("My UID is: %d. My GID is: %dn", current_uid, getgid());
    system("/usr/bin/id");
    if (setuid(0))
    {
        perror("setuid");
        return 1;
    }
    struct passwd *pw;
    pw = getpwnam("kiosk");
    if(pw->pw_uid==current_uid){
        system("/usr/bin/su - zabbix");
    }

    return 0;
}
