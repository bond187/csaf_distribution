// This file is Free Software under the MIT License
// without warranty, see README.md and LICENSES/MIT.txt for details.
//
// SPDX-License-Identifier: MIT
//
// SPDX-FileCopyrightText: 2023 German Federal Office for Information Security (BSI) <https://www.bsi.bund.de>
// Software-Engineering: 2023 Intevation GmbH <https://intevation.de>

// Package main implements the csaf_validator tool.
package main

import (
        "C"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/csaf-poc/csaf_distribution/v2/csaf"
	"github.com/csaf-poc/csaf_distribution/v2/util"
	//"github.com/jessevdk/go-flags"
)

//cgo cannot export structs to C without some workarounds. It looks like they are just for printing and 
//selecting options. Maybe we can hardcode the selection or pass it from python.
/*
type options struct {
	Version                bool     `long:"version" description:"Display version of the binary"`
	RemoteValidator        string   `long:"validator" description:"URL to validate documents remotely" value-name:"URL"`
	RemoteValidatorCache   string   `long:"validatorcache" description:"FILE to cache remote validations" value-name:"FILE"`
	RemoteValidatorPresets []string `long:"validatorpreset" description:"One or more presets to validate remotely" default:"mandatory"`
	Output                 string   `short:"o" long:"output" description:"If a remote validator was used, display AMOUNT ('all', 'important' or 'short') results" value-name:"AMOUNT"`
}
*/

//stuff in main needs to be commented out. It is not necessary for simple checking anyway.
func main() {
/*
	opts := new(options)

	parser := flags.NewParser(opts, flags.Default)
	parser.Usage = "[OPTIONS] files..."
	files, err := parser.Parse()
	errCheck(err)

	if opts.Version {
		fmt.Println(util.SemVersion)
		return
	}

	if len(files) == 0 {
		log.Println("No files given.")
		return
	}

	errCheck(run(opts, files))
*/
}


// run validates the given files. Some C types need to be used for export puposes. 
//export run
func run(file_c *C.char) int {
        file := C.GoString(file_c)
	//var validator csaf.RemoteValidator
        eval := util.NewPathEval()
        /*
	if opts.RemoteValidator != "" {
		validatorOptions := csaf.RemoteValidatorOptions{
			URL:     opts.RemoteValidator,
			Presets: opts.RemoteValidatorPresets,
			Cache:   opts.RemoteValidatorCache,
		}
		var err error
		if validator, err = validatorOptions.Open(); err != nil {
			return fmt.Errorf(
				"preparing remote validator failed: %w", err)
		}
		defer validator.Close()
	}
        */
	// Select amount level of output for remote validation.
        /*
	var printResult func(*csaf.RemoteValidationResult)
	switch opts.Output {
	case "all":
		printResult = printAll
	case "short":
		printResult = printShort
	case "important":
		printResult = printImportant
	case "":
		printResult = noPrint
	default:
		return fmt.Errorf("unknown output amount %q", opts.Output)
	}
        */

        //no loop. each call takes one file. Could easily use an array of file names again.
	//for _, file := range file_s {
                fault := 0
		// Check if the file name is valid.
		if !util.ConformingFileName(filepath.Base(file)) {
			fmt.Printf("%q is not a valid advisory name.\n", file)
		}
		doc, err := loadJSONFromFile(file)
		if err != nil {
			log.Printf("error: loading %q as JSON failed: %v\n", file, err)
			//continue
		}
		// Validate agsinst Schema.
		validationErrs, err := csaf.ValidateCSAF(doc)
		if err != nil {
			log.Printf("error: validating %q against schema failed: %v\n",
				file, err)
                }

		if len(validationErrs) > 0 {
			fmt.Printf("schema validation errors of %q\n", file)
			for _, vErr := range validationErrs {
				fmt.Printf("  * %s\n", vErr)
			}
                        fault = 1
		} else {
			fmt.Printf("%q passes the schema validation.\n", file)
		}

		// Check filename agains ID
		if err := util.IDMatchesFilename(eval, doc, filepath.Base(file)); err != nil {
			log.Printf("%s: %s.\n", file, err)
                        fault = 1
			//continue
                }


                /*
                Not currently using remote validator
		// Validate against remote validator.
		if validator != nil {
			rvr, err := validator.Validate(doc)
			if err != nil {
				return fmt.Errorf("remote validation of %q failed: %w",
					file, err)
			}
			printResult(rvr)
			var passes string
			if rvr.Valid {
				passes = "passes"
			} else {
				passes = "does not pass"
			}
			fmt.Printf("%q %s remote validation.\n", file, passes)
		}
                */
	//} end of for loop for if the parameter is an array of file names
        if fault == 0{
             return 0 //returns 0 instead of nil because C does not like the error type.
	} else {
	     return 1
	}
}

// noPrint suppresses the output of the validation result.
//func noPrint(*csaf.RemoteValidationResult) {}

// messageInstancePaths aggregates errors, warnings and infos by their
// message.

/*
type messageInstancePaths struct {
	message string
	paths   []string
}
*/

// messageInstancePathsList is a list for errors, warnings or infos.
//type messageInstancePathsList []messageInstancePaths

// addAll adds all errors, warnings or infos of a test.
/*
func (mipl *messageInstancePathsList) addAll(rtrs []csaf.RemoteTestResult) {
	for _, rtr := range rtrs {
		mipl.add(rtr)
	}
}
*/

// add adds a test result unless it is a duplicate.
/*
func (mipl *messageInstancePathsList) add(rtr csaf.RemoteTestResult) {
	for i := range *mipl {
		m := &(*mipl)[i]
		// Already have this message?
		if m.message == rtr.Message {
			for _, path := range m.paths {
				// Avoid dupes.
				if path == rtr.InstancePath {
					return
				}
			}
			m.paths = append(m.paths, rtr.InstancePath)
			return
		}
	}
	*mipl = append(*mipl, messageInstancePaths{
		message: rtr.Message,
		paths:   []string{rtr.InstancePath},
	})
}
*/

// print prints the details of the list to stdout if there are any.
/*
func (mipl messageInstancePathsList) print(info string) {
	if len(mipl) == 0 {
		return
	}
	fmt.Println(info)
	for i := range mipl {
		mip := &mipl[i]
		fmt.Printf("  message: %s\n", mip.message)
		fmt.Println("  instance path(s):")
		for _, path := range mip.paths {
			fmt.Printf("    %s\n", path)
		}
	}
}
*/

// printShort outputs the validation result in an aggregated version.
/*
func printShort(rvr *csaf.RemoteValidationResult) {

	var errors, warnings, infos messageInstancePathsList

	for i := range rvr.Tests {
		test := &rvr.Tests[i]
		errors.addAll(test.Error)
		warnings.addAll(test.Warning)
		infos.addAll(test.Info)
	}

	fmt.Printf("isValid: %t\n", rvr.Valid)
	errors.print("errors:")
	warnings.print("warnings:")
	infos.print("infos:")
}
*/

// printImportant displays only the test results which are really relevant.
/*
func printImportant(rvr *csaf.RemoteValidationResult) {
	printRemoteValidationResult(rvr, func(rt *csaf.RemoteTest) bool {
		return !rt.Valid ||
			len(rt.Info) > 0 || len(rt.Error) > 0 || len(rt.Warning) > 0
	})
}
*/

// printAll displays all test results.
/*
func printAll(rvr *csaf.RemoteValidationResult) {
	printRemoteValidationResult(rvr, func(*csaf.RemoteTest) bool {
		return true
	})
}
*/

// printInstanceAndMessages prints the message and the instance path of
// a test result.
/*
func printInstanceAndMessages(info string, me []csaf.RemoteTestResult) {
	if len(me) == 0 {
		return
	}
	fmt.Printf("  %s\n", info)
	for _, test := range me {
		fmt.Printf("    instance path: %s\n", test.InstancePath)
		fmt.Printf("    message: %s\n", test.Message)
	}
}
*/

// printRemoteValidationResult prints a filtered output of the remote validation result.
/*
func printRemoteValidationResult(
	rvr *csaf.RemoteValidationResult,
	accept func(*csaf.RemoteTest) bool,
) {

	fmt.Printf("isValid: %t\n", rvr.Valid)
	fmt.Println("tests:")
	nl := false
	for i := range rvr.Tests {
		test := &rvr.Tests[i]
		if !accept(test) {
			continue
		}
		if nl {
			fmt.Println()
		} else {
			nl = true
		}
		fmt.Printf("  name: %s\n", test.Name)
		fmt.Printf("  isValid: %t\n", test.Valid)
		printInstanceAndMessages("errors:", test.Error)
		printInstanceAndMessages("warnings:", test.Warning)
		printInstanceAndMessages("infos:", test.Info)
	}
}
*/

/*
func errCheck(err error) {
	if err != nil {
		if flags.WroteHelp(err) {
			os.Exit(0)
		}
		log.Fatalf("error: %v\n", err)
	}
}
*/

// loadJSONFromFile loads a JSON document from a file.
//export loadJSONFromFile
func loadJSONFromFile(fname string) (interface{}, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var doc any
	if err = json.NewDecoder(f).Decode(&doc); err != nil {
		return nil, err
	}
	return doc, err
}
