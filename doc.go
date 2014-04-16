// goXgo - Intro

/*
The basic idea for/behind goXgo is to be able to have external services
provide functionality that might not yet exist in go or would take too much
time to implement yourself - in an RPCish manner.

These services could be written in any other language and be running on
the same machine as the go code or be somewhere else.

The first idea for this came up when we talked about Natural Language
Processing (NLP) for which there are a couple of nice libraries - most
notably the NLTK (http://www.nltk.org).

Another thing we will need when dealing with NLP (/Pattern Matching/
Information Retreival/etc.) will be a solid scientific mathematics library:
I found some starts in go but nothing compared to numPy (http://www.numpy.org/).
R (http://www.r-project.org/) would come to mind too.

To start playing and testing NLP and statistical features i think it makes
a lot of sense to use stuff that's out there (and was written by linguists
or mathematicians specialized in scientific computing ;)

To network those services i chose
ZMQ(http://zguide.zeromq.org/page:all#-MQ-in-a-Hundred-Words).

Data is exchanged in JSON right now - MsgPack might be even more effective
(http://msgpack.org/).

The idea is to basically have a bunch of (static) functions that you want to
call from go:

1. You group those functions into Services

2. Give the service a name

3. And register the service with a "networking frontend"

On the python side this would look somthing like this

 from service_frontend import ZmqFrontend
 from service import Service
 from lib.static_nltk_wrappers import tokenize, stem

 NLTKService = Service( name = 'NLTK' )
 NLTKService.register_service_method( f = tokenize )
 NLTKService.register_service_method( f = stem )

 zmq_frontend = ZmqFrontend()
 zmq_frontend.register_service( NLTKService )
 zmq_frontend.start()


Try it:

First you need to install some stuff:

You need python2.7+ including dev headers, pip, libev4, libev-dev

Then you can install the needed python packages:

 ~/data/dev/go/src/goxgo [goxgo] $sudo pip install -r py_services/requirements.txt

To run the python service:

 $ python py_services/test_server.py
 Starting server. Listening on tcp://*:4243...
 Start serving.

Now the service runs and you can hit it from go.

 $ go test -v tests/goxgo_test.go
 === RUN TestGoXGoTokenize
 --- PASS: TestGoXGoTokenize (0.00 seconds)
         goxgo_test.go:34: tokenizeResponse:
                 Language: en
                 Tokens: [Give me a tokenized version of this body of text. Testing trying embodiment embodied]

 === RUN TestGoXGoStem
 --- PASS: TestGoXGoStem (0.00 seconds)
         goxgo_test.go:47: stemResponse:
                 Language: en
                 Stemmed Tokens: [give me a token version of this bodi of text. test tri embodi embodi]


Service/Function/Argument Naming/Mapping/Case convention

In python i registered a function tokenize:

 NLTKService = Service( name = 'NLTK' )
 NLTKService.register_service_method( f = tokenize )

In go i need to call it with this target:

 gxg.CallTarget { Services: []string{"NLTK/tokenize"}, Version: "0.1" }

The arguments for that function in python are:

 def tokenize( body, locale = 'en' ):

BUT in go my payload keys starts with uppercase letters:

 tokenizePayload := gxg.TokenizeRequest {
 	Target: &gxg.CallTarget {
		Services: []string{"NLTK/tokenize"},
		Version: "0.1" },
 	Body: "Give me a tokenized version of this body of text.",
 	Locale: "en",
 }

To keep case cannonical in both languages (in stuff exported in go...) the python
frontend lower-cases every parameter it finds. Unserialization from the JSON response
does not need something like that because I can rely on encoding/json/Unmarshal
to check the struct definition.

*/
package goxgo
