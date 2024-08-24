init_api:
	hz new --mod=xzdp --idl=idl/xzdp.thrift --customize_package=template/init.yaml --customize_layout_data_path=template/layout.yaml

update_api:
	hz update --mod=xzdp --idl=idl/xzdp.thrift --customize_package=template/package.yaml
	hz update --mod=xzdp --idl=idl/blog.thrift --customize_package=template/package.yaml
	hz update --mod=xzdp --idl=idl/user.thrift --customize_package=template/package.yaml
	hz update --mod=xzdp --idl=idl/shop.thrift --customize_package=template/package.yaml
	hz update --mod=xzdp --idl=idl/follow.thrift --customize_package=template/package.yaml
	hz update --mod=xzdp --idl=idl/image.thrift --customize_package=template/package.yaml
	hz update --mod=xzdp --idl=idl/blog_comment.thrift --customize_package=template/package.yaml
	hz update --mod=xzdp --idl=idl/message.thrift --customize_package=template/package.yaml
	hz update --mod=xzdp --idl=idl/voucher.thrift --customize_package=template/package.yaml

cleanhz:
	rm -rf biz *.go go.sum go.mod build.sh .hz conf output script .gitignore readme.md
