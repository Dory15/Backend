USE [Backend]
GO
/****** Object:  StoredProcedure [dbo].[insert_user]    Script Date: 10/04/2024 01:19:17 p. m. ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[insert_user]
	-- Add the parameters for the stored procedure here
	(
	@user VARCHAR (60),
	@email VARCHAR (60),
	@password VARCHAR (100),
	@phone VARCHAR(15)
	)
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

	BEGIN TRY 
	
		 DECLARE @trancount INT;
		 SET @trancount = @@TRANCOUNT;
			IF @trancount = 0
				BEGIN TRAN [tran_insert_user]; 
			ELSE
				SAVE TRANSACTION [tran_insert_user];


		IF EXISTS (
		SELECT 1 FROM [user] WHERE email = @email OR phone = @phone
		)
		BEGIN
			RAISERROR('Email or Phone has been registered', 16, 1)
			ROLLBACK TRAN
		END
		ELSE
		BEGIN
			INSERT INTO [user] ([user], email, [password], phone) VALUES (@user, @email, @password, @phone);
		END
		IF @trancount = 0
			COMMIT TRAN [tran_insert_user];
		END TRY
		BEGIN CATCH
		DECLARE @xstate INT

		SELECT @xstate = XACT_STATE();

		IF @xstate = -1
			ROLLBACK;
		IF @xstate = 1
			ROLLBACK TRANSACTION [tran_insert_user];
		END CATCH
	
END
