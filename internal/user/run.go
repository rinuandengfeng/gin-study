package user

import "context"

func Run(conf string) error {
	ctx := context.Background()

	service, err := newApiServer(conf)
	if err != nil {
		return err
	}

	return service.Run(ctx)
}
