#!/usr/bin/env python3

import argparse
import os
import subprocess


def convertAllInDir(mcfdir, outdir):
    for root, dirs, files in os.walk(mcfdir):
        for file in files:
            if file.endswith(".yml"):
                fullpath = os.path.join(root, file)
                print("Processing: ", fullpath)
                subprocess.call(
                    [
                        "pygeometa",
                        "generate-metadata",
                        "--mcf={:s}".format(fullpath),
                        "--schema=iso19139",
                        "--output={:s}/{:s}.xml".format(outdir, file[:-4]),
                    ]
                )


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--mcfdir", help="Path to MCF store")
    parser.add_argument("--outdir", help="Output directory", default=".")
    args = parser.parse_args()
    convertAllInDir(args.mcfdir, args.outdir)


if __name__ == "__main__":
    main()
