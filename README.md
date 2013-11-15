Coverage Author Extractor
=========================

import github.com/300brand/coverage to extract the author from the supplied URL. Keep in mind you will need to checkout
the **authors** branch in order to have the correct code.

Create a .json file with a structure that fits your needs, but something like this might be helpful:

```json
[
  {
    "Host": "wsj.com",
    "Paths": [
      "//meta[@name='article.author']/@content"
    ]
  }
]
```

Input
-----

Allow for the following flags:

* -xpath String representing XPath to use to search for the author. If not defined, attempt to use one defined in .json file
* -f Filename of aforementioned .json file. Default to "xpaths.json"
* Any non-flag arguments (accessible via flag.Args()) should be processed as URLs

Download the article and use the resulting []byte to pass into author.Search(). 

If an XPath is already defined for the URL's host and the -xpath flag is also defined, add the flag xpath to the 
beginning of the Paths array and use it in the author.Search() call.

Output
------

Author name and XPath used to find it.
