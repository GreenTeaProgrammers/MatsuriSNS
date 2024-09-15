import React, { useState } from "react";
import { useSelector, useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";
import { Typography, Container, Avatar, Button } from "@mui/material";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import { register } from "../store/authSlice";
import { RootState, AppDispatch } from "../store";
import {
  StyledTextField,
  AuthContainer,
  FormBox,
  SubmitButton,
} from "./AuthStyles";

const RegisterPage: React.FC = () => {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });
  const [confirmPassword, setConfirmPassword] = useState("");
  const [errors, setErrors] = useState<{ [key: string]: string }>({});

  const dispatch: AppDispatch = useDispatch();
  const navigate = useNavigate();

  // Accessing state from Redux store
  const authStatus = useSelector((state: RootState) => state.auth.status);
  const authError = useSelector((state: RootState) => state.auth.error);

  // Form validation
  const validateForm = (): boolean => {
    const newErrors: { [key: string]: string } = {};

    if (formData.password.length < 6) {
      newErrors.password = "パスワードは6文字以上である必要があります。";
    }
    if (formData.password !== confirmPassword) {
      newErrors.confirmPassword = "パスワードが一致しません。";
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (validateForm()) {
      dispatch(register(formData))
        .unwrap()
        .then(() => {
          navigate("/login");
        })
        .catch((err) => {
          if (err instanceof Error) {
            console.error("Registration error details:", err.message);
          } else {
            console.error("Unexpected error:", err);
          }
        });
    }
  };

  return (
    <Container component="main" maxWidth="xs">
      <AuthContainer elevation={6}>
        <Avatar sx={{ m: 1, bgcolor: "primary.main", width: 56, height: 56 }}>
          <LockOutlinedIcon fontSize="large" />
        </Avatar>
        <Typography component="h1" variant="h4" sx={{ mb: 3, fontWeight: 700 }}>
          ユーザー登録
        </Typography>
        {authStatus === "failed" && authError && (
          <Typography color="error">{authError}</Typography>
        )}
        <FormBox component="form" onSubmit={handleSubmit}>
          <StyledTextField
            required
            fullWidth
            id="username"
            label="ユーザー名"
            name="username"
            autoComplete="username"
            autoFocus
            value={formData.username}
            onChange={handleChange}
            variant="outlined"
          />
          <StyledTextField
            required
            fullWidth
            name="password"
            label="パスワード"
            type="password"
            id="password"
            autoComplete="new-password"
            value={formData.password}
            onChange={handleChange}
            error={!!errors.password}
            helperText={errors.password}
            variant="outlined"
          />
          <StyledTextField
            required
            fullWidth
            name="confirmPassword"
            label="パスワード（確認）"
            type="password"
            id="confirmPassword"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            error={!!errors.confirmPassword}
            helperText={errors.confirmPassword}
            variant="outlined"
          />
          <SubmitButton
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            disabled={authStatus === "loading"}
          >
            {authStatus === "loading" ? "登録中..." : "登録"}
          </SubmitButton>
          <Button
            fullWidth
            variant="text"
            onClick={() => navigate("/login")}
            sx={{ mt: 2, textTransform: "none" }}
          >
            すでにアカウントをお持ちの方はこちら
          </Button>
        </FormBox>
      </AuthContainer>
    </Container>
  );
};

export default RegisterPage;