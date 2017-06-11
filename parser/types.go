package parser

import "sync"
import "golang.org/x/exp/ebnf"

type Parser struct {
	lexFrames []lexicalFrame
	in        <-chan Token
}

// A lexecical frame stores all of the data associated with a
// single lexical token...
type lexicalFrame struct {
	token Token
	// each token gets it's own queue of parse frames
	workingQueue chan ParseFrame

	// each token also gets its own frame history
	// a frame history is a synchronized map...
	// from Productions -> ParseFrames
	record *FrameRecord
}

type ParseFrame struct {
	// has the link to the next frame in the chain...

	// has the link to the last frame in the chain...????

	// must know what production it is so it can walk to the next productions after it.
	prod *ebnf.Production

	// Must eventually be able to know if this frame is matched
	// by examining the input token when the time comes.
}

// FrameState represents  a map from Productions -> ParseFrames
type FrameRecord struct {
	mapper *sync.Map
}

// NewFrameState constructs a new map
func NewFrameRecord() *FrameRecord {
	return &FrameRecord{
		mapper: new(sync.Map),
	}
}

func (frameRecord *FrameRecord) Store(prod *ebnf.Production, frame *ParseFrame) {
	frameRecord.mapper.Store(prod, frame)
}

func (frameRecord *FrameRecord) Load(prod *ebnf.Production) (*ParseFrame, bool) {
	frame, ok := frameRecord.mapper.Load(prod)
	if !ok {
		return nil, false
	}

	return frame.(*ParseFrame), true
}

// Start is called when the parser is expected to begin execution
func (parser *Parser) Start() {
	go parser.Listen()
}

// Listen reads each token from the lexer and stores them in the
// token table
func (parser *Parser) Listen() {
	var next int = 0
	for token := range parser.in {
		parser.StoreTok(token, next)

		// fire off goroutine...
		go parser.Earley(next)

		next++
	}
}

func (parser *Parser) StoreTok(token Token, next int) {
	parser.tokens[next] = token
}

// this function performs an Earley parse for the given token at the given frame...
func (parser *Parser) Earley(index int) {

	myState := parser.lexFrames[index]
	for frame := range myState.workingQueue {
		// check if the frame has been see before...
		_, ok := myState.record.Load(frame.prod)
		if ok {
			// then this frame has already been explored.
			// nothing to do here
			// TODO double check what you do in this case...
			continue
		}

		// else, we need to expand the productions
		// and add them to the queue...

		parser.digestFrame(frame, index)
	}
}

// this function performs the type switch logic over
// the *ebnf.Production passed in
func (parser *Parser) digestFrame(frame *ParseFrame, index int) {
	rule := frame.prod.Expr
	switch rule.(type) {

	case *ebnf.Alternative:
		// digestAlternative
	case *ebnf.Sequence:
		// digestSequence
	case *ebnf.Repetition:
		// digestRepetition
	case *ebnf.Option:
		// digestOption
	case *ebnf.Group:
		// digestGroup
	case *ebnf.Range:
		// digestRange
	default:
		// panic
	}
}
