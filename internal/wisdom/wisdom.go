package wisdom

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Touch() string {
	const (
		padding   = 80
		copyright = "\u00a9 George Carlin"
	)

	// Select long quotations only
	//for i, n := 0, 0; i < len(georgeCarlinSay); i++ {
	//	if len(georgeCarlinSay[n]) <= padding {
	//		georgeCarlinSay = append(georgeCarlinSay[:n], georgeCarlinSay[n+1:]...)
	//	} else {
	//		n++
	//	}
	//}

	var (
		quotation = georgeCarlinSay[rand.Int()%len(georgeCarlinSay)]
		words     = strings.Split(quotation, " ")
		lines     []string
		buffer    = ""
	)

	flush := func() {
		lines = append(lines, buffer)
		buffer = ""
	}

	for i, w := range words {
		buffer = strings.Join([]string{buffer, w}, " ")
		if i >= len(words)-1 {
			flush()
		} else if len(buffer) >= padding {
			flush()
		}
	}

	return fmt.Sprintf(
		"%s\n%s%s",
		strings.Join(lines, "\n"),
		strings.Repeat(" ", padding-len(copyright)),
		copyright,
	)
}

var georgeCarlinSay = []string{
	"Don't sweat the petty things and don't pet the sweaty things.",
	"Inside every cynical person, there is a disappointed idealist.",
	"When you're born you get a ticket to the freak show. When you're born in America, you get a front row seat.",
	"Have you ever noticed that anybody driving slower than you is an idiot, and anyone going faster than you is a maniac?",
	"Just cause you got the monkey off your back doesn't mean the circus has left town.",
	"Weather forecast for tonight: dark.",
	"Electricity is really just organized lightning.",
	"Some people see things that are and ask, Why? Some people dream of things that never were and ask, Why not? Some people have to go to work and don't have time for all that.",
	"May the forces of evil become confused on the way to your house.",
	"There are nights when the wolves are silent and only the moon howls.",
	"Most people work just hard enough not to get fired and get paid just enough money not to quit.",
	"Well, if crime fighters fight crime and fire fighters fight fire, what do freedom fighters fight? They never mention that part to us, do they?",
	"Frisbeetarianism is the belief that when you die, your soul goes up on the roof and gets stuck.",
	"If you can't beat them, arrange to have them beaten.",
	"There's no present. There's only the immediate future and the recent past.",
	"People who say they don't care what people think are usually desperate to have people think they don't care what people think.",
	"The very existence of flame-throwers proves that some time, somewhere, someone said to themselves, You know, I want to set those people over there on fire, but I'm just not close enough to get the job done.",
	"One tequila, two tequila, three tequila, floor.",
	"The other night I ate at a real nice family restaurant. Every table had an argument going.",
	"By and large, language is a tool for concealing the truth.",
	"Death is caused by swallowing small amounts of saliva over a long period of time.",
	"I'm completely in favor of the separation of Church and State. My idea is that these two institutions screw us up enough on their own, so both of them together is certain death.",
	"In comic strips, the person on the left always speaks first.",
	"If it's true that our species is alone in the universe, then I'd have to say the universe aimed rather low and settled for very little.",
	"I went to a bookstore and asked the saleswoman, 'Where's the self-help section?' She said if she told me, it would defeat the purpose.",
	"Religion is just mind control.",
	"The main reason Santa is so jolly is because he knows where all the bad girls live.",
	"Atheism is a non-prophet organization.",
	"'I am' is reportedly the shortest sentence in the English language. Could it be that 'I do' is the longest sentence?",
	"I recently went to a new doctor and noticed he was located in something called the Professional Building. I felt better right away.",
	"Not only do I not know what's going on, I wouldn't know what to do about it if I did.",
	"The reason I talk to myself is that I'm the only one whose answers I accept.",
	"I was thinking about how people seem to read the Bible a whole lot more as they get older; then it dawned on me - they're cramming for their final exam.",
	"I think people should be allowed to do anything they want. We haven't tried that for a while. Maybe this time it'll work.",
	"Dusting is a good example of the futility of trying to put things right. As soon as you dust, the fact of your next dusting has already been established.",
	"What does it mean to pre-board? Do you get on before you get on?",
	"Always do whatever's next.",
	"I'm not concerned about all hell breaking loose, but that a PART of hell will break loose... it'll be much harder to detect.",
	"You know an odd feeling? Sitting on the toilet eating a chocolate candy bar.",
	"One can never know for sure what a deserted area looks like.",
	"I have as much authority as the Pope, I just don't have as many people who believe it.",
	"I would never want to be a member of a group whose symbol was a guy nailed to two pieces of wood.",
	"When Thomas Edison worked late into the night on the electric light, he had to do it by gas lamp or candle. I'm sure it made the work seem that much more urgent.",
	"Think off-center.",
	"I think it would be interesting if old people got anti-Alzheimer's disease where they slowly began to recover other people's lost memories.",
	"At a formal dinner party, the person nearest death should always be seated closest to the bathroom.",
	"I'm always relieved when someone is delivering a eulogy and I realize I'm listening to it.",
	"When someone is impatient and says, 'I haven't got all day,' I always wonder, How can that be? How can you not have all day?",
	"If we could just find out who's in charge, we could kill him.",
	"When you step on the brakes your life is in your foot's hands.",
	"The status quo sucks.",
	"Standing ovations have become far too commonplace. What we need are ovations where the audience members all punch and kick one another.",
	"You know the good part about all those executions in Texas? Fewer Texans.",
}
