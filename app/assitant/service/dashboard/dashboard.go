package dashboard

import (
	"context"

	"github.com/yira97/cnworld/app/assitant/service/application"
	"github.com/yira97/cnworld/app/assitant/service/user"
	"github.com/yira97/cnworld/app/assitant/storage"
	"go.mongodb.org/mongo-driver/bson"
)

type DashBoardView struct {
	VU  storage.UserView
	VAs []storage.ApplicationView
}

type DashBorad struct {
	U  user.User
	As []application.Application
}

func (d *DashBorad) UserSync(ctx context.Context, info user.UserIdentifiableInfo) error {
	err := d.U.Sync(ctx, info)
	if err != nil {
		return err
	}
	return nil
}

// 确保运行前运行了 UserSync
func (d *DashBorad) ApplicationSync(ctx context.Context) error {
	dumbApp := storage.ApplicationStorage{}

	// filter写法参考
	// https://stackoverflow.com/questions/37231760/golang-mongodb-having-issues-using-in-to-find-all-elements-with-one-string-in
	memberF := bson.E{
		Key:   storage.ApplicationStorage_Member,
		Value: bson.M{"$in": d.U.V.UID},
	}
	filter := dumbApp.DefaultOrderedFilter(memberF)

	cursor, err := storage.ApplicationStorage{}.Collection().Find(ctx, filter)
	if err != nil {
		return err
	}

	apps := []storage.ApplicationStorage{}
	err = cursor.All(ctx, apps)
	if err != nil {
		return err
	}

	for _, app := range apps {
		d.As = append(d.As, application.Application{
			V: app.View(),
			M: &app,
		})
	}
	return nil
}

func (d *DashBorad) Sync(ctx context.Context, info user.UserIdentifiableInfo) error {
	err := d.UserSync(ctx, info)
	if err != nil {
		return err
	}
	err = d.ApplicationSync(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (d DashBorad) View() DashBoardView {
	v := DashBoardView{}
	v.VU = d.U.V
	for _, a := range d.As {
		v.VAs = append(v.VAs, a.V)
	}
	return v
}

func GetDashBoard(ctx context.Context, info user.UserIdentifiableInfo) (*DashBoardView, error) {
	d := DashBorad{}
	err := d.Sync(ctx, info)
	if err != nil {
		return nil, err
	}
	v := d.View()
	return &v, nil
}
