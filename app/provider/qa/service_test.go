package qa

import (
	"bbs/app/provider/user"
	"bbs/test"
	"context"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/provider/cache"
	"github.com/gohade/hade/framework/provider/config"
	"github.com/gohade/hade/framework/provider/log"
	"github.com/gohade/hade/framework/provider/orm"
	"github.com/gohade/hade/framework/provider/redis"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test_QA(t *testing.T) {
	container := test.InitBaseContainer()
	container.Bind(&config.HadeConfigProvider{})
	container.Bind(&log.HadeLogServiceProvider{})
	container.Bind(&orm.GormProvider{})
	container.Bind(&redis.RedisProvider{})
	container.Bind(&cache.HadeCacheProvider{})
	container.Bind(&user.UserProvider{})

	//userService := container.MustMake(user.UserKey).(user.Service)
	ormService := container.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := ormService.GetDB()
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&Question{}, &Answer{}); err != nil {
		t.Fatal(err)
	}

	if err := db.Exec("delete from answers where 1").Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Exec("delete from questions where 1").Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Exec("delete from users where 1").Error; err != nil {
		t.Fatal(err)
	}

	tmp, err := NewQaService(container)
	if err != nil {
		t.Fatal(err)
	}
	qaService := tmp.(Service)

	user1 := &user.User{
		UserName:  "user1",
		Email:     "user1@gmail.com",
		CreatedAt: time.Now(),
	}
	user2 := &user.User{
		UserName:  "user2",
		Email:     "user2@gmail.com",
		CreatedAt: time.Now(),
	}

	ctx := context.Background()

	db.Create(user1)

	db.Create(user2)

	Convey("正常流程", t, func() {
		var question1 = &Question{
			Title:     "question1",
			Context:   "this is context",
			AnswerNum: 0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		var question2 = &Question{
			Title:     "question2",
			Context:   "this is context",
			AnswerNum: 0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		{
			question1.AuthorID = user1.ID
			err := qaService.PostQuestion(ctx, question1)
			So(err, ShouldBeNil)

			question1, err = qaService.GetQuestion(ctx, question1.ID)
			So(err, ShouldBeNil)
		}

		{
			question2.AuthorID = user2.ID
			err := qaService.PostQuestion(ctx, question2)
			So(err, ShouldBeNil)

			question2, err = qaService.GetQuestion(ctx, question2.ID)
			So(err, ShouldBeNil)
		}

		{
			q, err := qaService.GetQuestion(ctx, question1.ID)
			So(err, ShouldBeNil)
			So(q.Title, ShouldEqual, question1.Title)
		}

		{
			qs, err := qaService.GetQuestions(ctx, &Pager{
				Start: 0,
				Size:  10,
			})
			So(err, ShouldBeNil)
			So(qs, ShouldNotBeNil)
			So(len(qs), ShouldEqual, 2)
		}

		{
			err := qaService.QuestionLoadAuthor(ctx, question1)
			So(err, ShouldBeNil)
			So(question1.Author.ID, ShouldEqual, user1.ID)
		}

		{
			questions, err := qaService.GetQuestions(ctx, &Pager{Start: 0, Size: 10})
			So(err, ShouldBeNil)
			err = qaService.QuestionsLoadAuthor(ctx, &questions)
			So(err, ShouldBeNil)
			So(len(questions), ShouldEqual, 2)
			So(questions[1].Author.ID, ShouldEqual, user2.ID)
		}

		answer1 := &Answer{
			QuestionID: question1.ID,
			Content:    "answer context",
			AuthorID:   user2.ID,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		{
			err := qaService.PostAnswer(ctx, answer1)
			So(err, ShouldBeNil)
		}

		{
			err := qaService.QuestionLoadAnswers(ctx, question1)
			So(err, ShouldBeNil)
			So(question1.AnswerNum, ShouldEqual, 1)
			So(question1.Answers, ShouldNotBeNil)
			So(len(question1.Answers), ShouldEqual, 1)
		}

		{
			question1.Answers = nil
			qs := []*Question{question1, question2}
			err := qaService.QuestionsLoadAnswers(ctx, &qs)
			So(err, ShouldBeNil)
			So(qs[0].Answers, ShouldNotBeNil)
			So(len(qs[0].Answers), ShouldEqual, 1)
		}

		{
			an, err := qaService.GetAnswer(ctx, answer1.ID)
			So(err, ShouldBeNil)
			So(an, ShouldNotBeNil)
			So(an.Content, ShouldEqual, answer1.Content)
		}

		{
			question1.Title = "question1 content update"
			err := qaService.UpdateQuestion(ctx, question1)
			So(err, ShouldBeNil)
		}

		answer2 := &Answer{
			QuestionID: question2.ID,
			Content:    "answer2 content",
			AuthorID:   user1.ID,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		{
			err := qaService.PostAnswer(ctx, answer2)
			So(err, ShouldBeNil)
		}

		{
			err := qaService.DeleteAnswer(ctx, answer2.ID)
			So(err, ShouldBeNil)
		}

		{
			err := qaService.DeleteAnswer(ctx, answer1.ID)
			So(err, ShouldBeNil)

			err = qaService.DeleteQuestion(ctx, question1.ID)
			So(err, ShouldBeNil)
		}

	})

}
