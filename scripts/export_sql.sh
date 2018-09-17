#!/bin/bash

sqlite3 data/data.db ".dump" > tmp/db.sql
