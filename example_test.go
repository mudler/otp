// Copyright (C) 2019 Michael J. Fromberger. All Rights Reserved.

package otp_test

import (
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/creachadair/otp"
)

func fixedTime(z uint64) func() uint64 { return func() uint64 { return z } }

func Example() {
	cfg := otp.Config{
		Hash:   sha256.New, // default is sha1.New
		Digits: 8,          // default is 6
		// By default, time-based OTP generation uses time.Now.  You can plug in
		// your own function to control how time steps are generated.
		// This example uses a fixed time step so the output will be consistent.
		TimeStep: fixedTime(1),
	}

	// 2FA setup tools often present the shared secret as a base32 string.
	// ParseKey decodes this format.
	if err := cfg.ParseKey("MFYH A3DF EB2G C4TU"); err != nil {
		log.Fatalf("Parsing key: %v", err)
	}

	fmt.Println("HOTP", 0, cfg.HOTP(0))
	fmt.Println("HOTP", 1, cfg.HOTP(1))
	fmt.Println()
	fmt.Println("TOTP", cfg.TOTP())
	// Output:
	// HOTP 0 59590364
	// HOTP 1 86761489
	//
	// TOTP 86761489
}

func Example2() {
	cfg := otp.Config{
		Hash:    sha256.New, // default is sha1.New
		Digits:  44,         // default is 6
		NoTrunc: true,
		// By default, time-based OTP generation uses time.Now.  You can plug in
		// your own function to control how time steps are generated.
		// This example uses a fixed time step so the output will be consistent.
		TimeStep: fixedTime(1),
	}

	// 2FA setup tools often present the shared secret as a base32 string.
	// ParseKey decodes this format.
	if err := cfg.ParseKey("MFYH A3DF EB2G C4TU"); err != nil {
		log.Fatalf("Parsing key: %v", err)
	}

	fmt.Println("HOTP", 0, cfg.HOTP(0))
	fmt.Println("HOTP", 1, cfg.HOTP(1))
	fmt.Println()
	fmt.Println("TOTP", cfg.TOTP())
	// Output:
	// HOTP 0 EIiZPygQ3ArK2KObo3ILwr026IWbRkvS3zm/413yHPM=
	// HOTP 1 6WKBpzdqZy4jIijvJhHK2LiGDjoFkaL8JUkW7ASd8po=
	//
	// TOTP 6WKBpzdqZy4jIijvJhHK2LiGDjoFkaL8JUkW7ASd8po=
}

func ExampleConfig_customFormat() {
	// Use settings compatible with Steam Guard: 5 characters and a custom alphabet.
	cfg := otp.Config{
		Digits:   5,
		Format:   otp.FormatAlphabet("23456789BCDFGHJKMNPQRTVWXY"),
		TimeStep: fixedTime(9876543210),
	}
	if err := cfg.ParseKey("CQKQ QEQR AAR7 77X5"); err != nil {
		log.Fatalf("Parsing key: %v", err)
	}

	fmt.Println(cfg.TOTP())
	// Output:
	// FKNK3
}
