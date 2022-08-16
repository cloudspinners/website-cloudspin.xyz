#!/usr/bin/env sh

set -euo pipefail

fail()
{
  >&2 echo "FAIL: $*"
  exit 1
}


cd site


for arg in "$@"
do
	case "$arg" in
	preview)
		echo "Including unpublished posts"
		jekyll serve --host=0.0.0.0 --watch --unpublished
		;;
	plan)
		echo "Only including published posts"
		jekyll serve --host=0.0.0.0 --watch
		;;
	publish)
		[ -n "${WEBSITE_S3_BUCKET-}" ] || fail "WEBSITE_S3_BUCKET not set"
		[ -n "${WEBSITE_AWS_PROFILE-}" ] || fail "WEBSITE_AWS_PROFILE not set"
		jekyll build
		aws s3 \
			--profile ${WEBSITE_AWS_PROFILE} \
			sync \
			--acl public-read \
			--exact-timestamps \
			--delete \
			./_site/ \
			s3://${WEBSITE_S3_BUCKET}/
		;;
	build)
		jekyll build
		;;
	help)
		echo "Usage: go preview|plan|publish|build"
		exit 0
		;;
	*)
		echo "Unknown command $arg"
		echo "Usage: go preview|plan|publish|build"
		exit 1
		;;
	esac
done
