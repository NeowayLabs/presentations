#!/usr/bin/env nash

fn startConsumer(dchanhost) {
	rfork upm {
		mount -t 9p $dchanhost /mnt // HL1
		./worker -i /mnt/projX/input -o /mnt/projX/output // HL2
		umount /mnt // HL1
	}
}

fn startProducer(dburi, dchanhost) {
	rfork upm {
		mount -t 9p $dchanhost /mnt // HL3
		./geninput -dburi $dburi -o /mnt/projX/input // HL4
		umount /mnt // HL3
	}
}

# END OMIT

fn main() {
	if len($ARGS) == "1" {
		printf "Usage: %s [producer consumer]\n"

		return
	}

	cmd = $ARGS[1]

	if $cmd == "producer" {
		startProducer($DBURI, $DCHANHOST)
	} else if $cmd == "consumer" {
		startConsumer($DCHANHOST)
	} else {
		printf "Invalid option: %s\n" $cmd
	}
}

# END OMIT
