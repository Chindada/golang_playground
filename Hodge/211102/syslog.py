import random
import logging
import datetime
import sys
from logging.handlers import SysLogHandler


device_hosts = ['127.0.0.1', '169.254.0.0']
tags = ['robotframework: ']
all_facility = [0, 1, 2, 3, 4]
all_serverity = [50, 40, 30]
all_message = ['syslog test', 'e2e test', 'test']

host = sys.argv[1]
port = 567


def syslogs_sender():
    now = datetime.datetime.now().strftime('%b %d %H:%M:%S')
    for ip in device_hosts:
        for custom_facility in all_facility:
            logger = logging.getLogger()
            syslog = SysLogHandler(
                address=(host, port), facility=custom_facility)
            logger.addHandler(syslog)
            random_tag = random.choice(tags)
            fields = {
                'host_field': ip,
                'date_field': now,
                'tag_field': random_tag,
            }
            format = logging.Formatter(
                '%(date_field)s %(host_field)s {0}: %(message)s'.format(random_tag))
            syslog.setFormatter(format)
            for level in all_serverity:
                logger.log(level, random.choice(
                    all_message), extra=fields)
            logger.removeHandler(syslog)
            syslog.close()


if __name__ == '__main__':
    syslogs_sender()
