/*
Package compflag generates bash completions on-the-fly from stdlib flag values.

Usage - just call compflag.Complete() somewhere before actual app logic, best point is right at the start:

	func main() {
	  if compflag.Complete() {
	    os.Exit(0)
	  }

	  flag.Parse()

	  // other startup logic...
	}

Please note, that you need to exit app if any completion happened.

Build your app, put binary somewhere in your "PATH", then run:

	complete -C %your-binary-name% %your-binary-name%

Now enter "%your-binary-name%", and hit tab twice )
*/
package compflag
